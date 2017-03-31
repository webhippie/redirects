package source

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// SetSource initializes the source URL.
func SetSource() gin.HandlerFunc {
	s := source{}

	return func(c *gin.Context) {
		s.applyToContext(c)
	}
}

// Get returns the source for redirects.
func Get(c *gin.Context) *url.URL {
	v1, ok := c.Get("source")

	if !ok {
		return nil
	}

	v2, ok := v1.(*url.URL)

	if !ok {
		return nil
	}

	return v2
}

// source represents just a helper to gather the source.
type source struct{}

// applyToContext applies the value to the context.
func (s *source) applyToContext(c *gin.Context) {
	value := c.Request.URL

	value.Scheme = s.resolveScheme(c.Request)
	value.Host = s.resolveDomain(c.Request)

	c.Set("source", value)
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
func (s *source) resolveDomain(r *http.Request) (host string) {
	switch {
	case r.Header.Get("X-Forwarded-For") != "":
		return r.Header.Get("X-Forwarded-For")
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
