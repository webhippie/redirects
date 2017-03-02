package store

import (
	"github.com/tboerger/redirects/model"
	// "golang.org/x/crypto/acme/autocert"
)

// Store implements all required data-layer functions for Redirects.
type Store interface {
	// GetRedirects retrieves all redirects from the store.
	GetRedirects() (*model.Redirects, error)

	// CreateRedirect creates a redirect on the store.
	CreateRedirect(*model.Redirect) error

	// UpdateRedirect updates a redirect on the store.
	UpdateRedirect(*model.Redirect) error

	// DeleteRedirect deletes a redirect from the store.
	DeleteRedirect(*model.Redirect) error

	// GetRedirect retrieves a specific redirect from the store.
	GetRedirect(string) (*model.Redirect, error)

	// CertCache returns the cert cache implementation.
	// CertCache() autocert.Cache
}
