package zookeeper

import (
	"encoding/json"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/webhippie/redirects/pkg/model"
	"github.com/webhippie/redirects/pkg/store"
)

// GetRedirects retrieves all redirects from the Zookeeper store.
func (db *data) GetRedirects() ([]*model.Redirect, error) {
	return db.load()
}

// GetRedirect retrieves a specific redirect from the Zookeeper store.
func (db *data) GetRedirect(id string) (*model.Redirect, error) {
	redirects, err := db.load()

	if err != nil {
		return nil, err
	}

	for _, record := range redirects {
		if record.ID == id {
			return record, nil
		}
	}

	return nil, store.ErrRedirectNotFound
}

// DeleteRedirect deletes a redirect from the Zookeeper store.
func (db *data) DeleteRedirect(id string) error {
	// TODO: Add distributed locking

	if ok, _ := db.store.Exists(db.key(id)); !ok {
		return store.ErrRedirectNotFound
	}

	return db.store.Delete(db.key(id))
}

// UpdateRedirect updates a redirect on the Zookeeper store.
func (db *data) UpdateRedirect(update *model.Redirect) error {
	// TODO: Add distributed locking

	if ok, _ := db.store.Exists(db.key(update.ID)); !ok {
		return store.ErrRedirectNotFound
	}

	redirects, err := db.load()

	if err != nil {
		return err
	}

	for _, record := range redirects {
		if record.ID != update.ID && record.Source == update.Source {
			return store.ErrRedirectSourceExists
		}
	}

	bytes, err := json.Marshal(update)

	if err != nil {
		return fmt.Errorf("failed to marshal record: %w", err)
	}

	return db.store.Put(
		db.key(update.ID),
		bytes,
		nil,
	)
}

// CreateRedirect creates a redirect on the Zookeeper store.
func (db *data) CreateRedirect(create *model.Redirect) error {
	id := uuid.NewV4().String()
	redirects, err := db.load()

	if err != nil {
		return err
	}

	for _, record := range redirects {
		if record.Source == create.Source {
			return store.ErrRedirectSourceExists
		}
	}

	create.ID = id
	bytes, err := json.Marshal(create)

	if err != nil {
		return fmt.Errorf("failed to marshal record: %w", err)
	}

	return db.store.Put(
		db.key(id),
		bytes,
		nil,
	)
}
