package yaml

import (
	uuid "github.com/satori/go.uuid"
	"github.com/webhippie/redirects/pkg/model"
	"github.com/webhippie/redirects/pkg/store"
)

// GetRedirects retrieves all redirects from the YAML store.
func (db *data) GetRedirects() ([]*model.Redirect, error) {
	root, err := db.load()

	if err != nil {
		return nil, err
	}

	return root.Redirects, nil
}

// GetRedirect retrieves a specific redirect from the YAML store.
func (db *data) GetRedirect(id string) (*model.Redirect, error) {
	root, err := db.load()

	if err != nil {
		return nil, err
	}

	for _, record := range root.Redirects {
		if record.ID == id {
			return record, nil
		}
	}

	return nil, store.ErrRedirectNotFound
}

// DeleteRedirect deletes a redirect from the YAML store.
func (db *data) DeleteRedirect(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	root, err := db.load()

	if err != nil {
		return err
	}

	for row, record := range root.Redirects {
		if record.ID == id {
			root.Redirects = append(
				root.Redirects[:row],
				root.Redirects[row+1:]...,
			)

			return db.write(root)
		}
	}

	return store.ErrRedirectNotFound
}

// UpdateRedirect updates a redirect on the YAML store.
func (db *data) UpdateRedirect(update *model.Redirect) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	root, err := db.load()

	if err != nil {
		return err
	}

	for row, record := range root.Redirects {
		if record.ID == update.ID {
			root.Redirects[row] = update
			return db.write(root)
		}
	}

	return store.ErrRedirectNotFound
}

// CreateRedirect creates a redirect on the YAML store.
func (db *data) CreateRedirect(create *model.Redirect) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	root, err := db.load()

	if err != nil {
		return err
	}

	for _, record := range root.Redirects {
		if record.Source == create.Source {
			return store.ErrRedirectSourceExists
		}
	}

	create.ID = uuid.NewV4().String()

	root.Redirects = append(
		root.Redirects,
		create,
	)

	return db.write(root)
}
