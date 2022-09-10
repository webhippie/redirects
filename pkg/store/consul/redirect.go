package consul

import (
	"context"
	"encoding/json"
	"fmt"

	valkeyrieStore "github.com/kvtools/valkeyrie/store"
	uuid "github.com/satori/go.uuid"
	"github.com/webhippie/redirects/pkg/model"
	"github.com/webhippie/redirects/pkg/store"
)

// GetRedirects retrieves all redirects from the Consul store.
func (db *data) GetRedirects(ctx context.Context) ([]*model.Redirect, error) {
	return db.load(ctx)
}

// GetRedirect retrieves a specific redirect from the Consul store.
func (db *data) GetRedirect(ctx context.Context, id string) (*model.Redirect, error) {
	redirects, err := db.load(ctx)

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

// DeleteRedirect deletes a redirect from the Consul store.
func (db *data) DeleteRedirect(ctx context.Context, id string) error {
	// TODO: Add distributed locking

	if ok, _ := db.store.Exists(ctx, db.key(id), &valkeyrieStore.ReadOptions{}); !ok {
		return store.ErrRedirectNotFound
	}

	return db.store.Delete(ctx, db.key(id))
}

// UpdateRedirect updates a redirect on the Consul store.
func (db *data) UpdateRedirect(ctx context.Context, update *model.Redirect) error {
	// TODO: Add distributed locking

	if ok, _ := db.store.Exists(ctx, db.key(update.ID), &valkeyrieStore.ReadOptions{}); !ok {
		return store.ErrRedirectNotFound
	}

	redirects, err := db.load(ctx)

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
		ctx,
		db.key(update.ID),
		bytes,
		nil,
	)
}

// CreateRedirect creates a redirect on the Consul store.
func (db *data) CreateRedirect(ctx context.Context, create *model.Redirect) error {
	id := uuid.NewV4().String()
	redirects, err := db.load(ctx)

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
		ctx,
		db.key(id),
		bytes,
		nil,
	)
}
