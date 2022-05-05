package source

import (
	"context"
	"net/http"
	"net/url"
	"strings"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	// ContextKey defines the context key used to store within the context.
	ContextKey = contextKey("source")
)

// Source writes the parsed source URL into the context.
func Source(next http.Handler) http.Handler {
	s := source{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(
			w,
			r.WithContext(
				s.fromRequest(r),
			),
		)
	})
}

// Get returns the source for redirects from context.
func Get(ctx context.Context) *url.URL {
	val, ok := ctx.Value(ContextKey).(*url.URL)

	if !ok {
		return nil
	}

	return val
}

// source represents just a helper to gather the source.
type source struct{}

// applyToContext applies the value to the context.
func (s *source) fromRequest(r *http.Request) context.Context {
	value := r.URL

	value.Scheme = s.resolveScheme(r)
	value.Host = s.resolveDomain(r)

	return context.WithValue(
		r.Context(),
		ContextKey,
		value,
	)
}

// resolveScheme retrieves the requested scheme.
func (s *source) resolveScheme(r *http.Request) string {
	switch {
	case r.Header.Get("X-Forwarded-Proto") == "https":
		return "https"
	case r.URL.Scheme == "https":
		return "https"
	case r.TLS != nil:
		return "https"
	case strings.HasPrefix(r.Proto, "HTTPS"):
		return "https"
	default:
		return "http"
	}
}

// resolveDomain retrieves the requested domain.
func (s *source) resolveDomain(r *http.Request) string {
	switch {
	case r.Header.Get("X-Forwarded-Host") != "":
		return r.Header.Get("X-Forwarded-Host")
	case r.Header.Get("X-Host") != "":
		return r.Header.Get("X-Host")
	case r.Host != "":
		return r.Host
	case r.URL.Host != "":
		return r.URL.Host
	default:
		return "localhost:8080"
	}
}
