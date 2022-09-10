package store

import (
	"context"
	"fmt"

	"github.com/webhippie/redirects/pkg/model"
)

var (
	// ErrRedirectNotFound gets returned if a redirect can't be found on the store.
	ErrRedirectNotFound = fmt.Errorf("failed to find a redirect")

	// ErrRedirectSourceExists gets returned if a redirect s source already exists.
	ErrRedirectSourceExists = fmt.Errorf("source already exists")
)

// Store implements all required data-layer functions for Redirects.
type Store interface {
	// Name simply returns the name of the store implementation.
	Name() string

	// Config just returns a simple configuration explanation.
	Config() string

	// GetRedirects retrieves all redirects from the store.
	GetRedirects(context.Context) ([]*model.Redirect, error)

	// GetRedirect retrieves a specific redirect from the store.
	GetRedirect(context.Context, string) (*model.Redirect, error)

	// DeleteRedirect deletes a redirect from the store.
	DeleteRedirect(context.Context, string) error

	// CreateRedirect creates a redirect on the store.
	CreateRedirect(context.Context, *model.Redirect) error

	// UpdateRedirect updates a redirect on the store.
	UpdateRedirect(context.Context, *model.Redirect) error
}
