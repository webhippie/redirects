package store

import (
	"github.com/tboerger/redirects/model"
	"golang.org/x/net/context"
)

// GetRedirects retrieves all redirects from the store.
func GetRedirects(c context.Context) ([]*model.Redirect, error) {
	return FromContext(c).GetRedirects()
}

// GetRedirect retrieves a specific redirect from the store.
func GetRedirect(c context.Context, id string) (*model.Redirect, error) {
	return FromContext(c).GetRedirect(id)
}

// DeleteRedirect deletes a redirect from the store.
func DeleteRedirect(c context.Context, id string) error {
	return FromContext(c).DeleteRedirect(id)
}

// CreateRedirect creates a redirect on the store.
func CreateRedirect(c context.Context, record *model.Redirect) error {
	return FromContext(c).CreateRedirect(record)
}

// UpdateRedirect updates a redirect on the store.
func UpdateRedirect(c context.Context, record *model.Redirect) error {
	return FromContext(c).UpdateRedirect(record)
}
