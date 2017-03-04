package store

import (
	"fmt"
	"github.com/tboerger/redirects/model"
	// "golang.org/x/crypto/acme/autocert"
)

var (
	// ErrRedirectNotFound gets returned if a redirect can't be found on the store.
	ErrRedirectNotFound = fmt.Errorf("Failed to find a redirect")
)

// Store implements all required data-layer functions for Redirects.
type Store interface {
	// GetRedirects retrieves all redirects from the store.
	GetRedirects() ([]*model.Redirect, error)

	// GetRedirect retrieves a specific redirect from the store.
	GetRedirect(int) (*model.Redirect, error)

	// DeleteRedirect deletes a redirect from the store.
	DeleteRedirect(int) error

	// CreateRedirect creates a redirect on the store.
	CreateRedirect(*model.Redirect) error

	// UpdateRedirect updates a redirect on the store.
	UpdateRedirect(*model.Redirect) error

	// CertCache returns the cert cache implementation.
	// CertCache() autocert.Cache
}
