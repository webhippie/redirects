package yaml

import (
	"github.com/tboerger/redirects/model"
)

// GetRedirects retrieves all redirects from the YAML store.
func (db *data) GetRedirects() (*model.Redirects, error) {
	return &model.Redirects{}, nil
}

// CreateRedirect creates a redirect on the YAML store.
func (db *data) CreateRedirect(record *model.Redirect) error {
	return nil
}

// UpdateRedirect updates a redirect on the YAML store.
func (db *data) UpdateRedirect(record *model.Redirect) error {
	return nil
}

// DeleteRedirect deletes a redirect from the YAML store.
func (db *data) DeleteRedirect(record *model.Redirect) error {
	return nil
}

// GetRedirect retrieves a specific redirect from the YAML store.
func (db *data) GetRedirect(id string) (*model.Redirect, error) {
	return &model.Redirect{}, nil
}
