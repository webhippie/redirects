package etcd

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/tboerger/redirects/model"
	"github.com/tboerger/redirects/store"
	"path"
	"time"
)

// GetRedirects retrieves all redirects from the Etcd store.
func (db *data) GetRedirects() ([]*model.Redirect, error) {
	return db.load()
}

// GetRedirect retrieves a specific redirect from the Etcd store.
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

// DeleteRedirect deletes a redirect from the Etcd store.
func (db *data) DeleteRedirect(id string) error {
	// TODO: Add distributed locking

	if ok, _ := db.store.Exists(db.key(id)); !ok {
		return store.ErrRedirectNotFound
	}

	return db.store.Delete(db.key(id))
}

// UpdateRedirect updates a redirect on the Etcd store.
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
		return fmt.Errorf("Failed to marshal record. %s", err)
	}

	return db.store.Put(
		db.key(update.ID),
		bytes,
		nil,
	)
}

// CreateRedirect creates a redirect on the Etcd store.
func (db *data) CreateRedirect(create *model.Redirect) error {
	id := db.nextID()
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
		return fmt.Errorf("Failed to marshal record. %s", err)
	}

	return db.store.Put(
		db.key(id),
		bytes,
		nil,
	)
}

// nextID tries to generate a new unique ID.
func (db *data) nextID() string {
	return fmt.Sprintf(
		"%x",
		md5.Sum([]byte(
			fmt.Sprintf("%v", time.Now().Unix()),
		)),
	)
}

// key generates the new key including the prefix.
func (db *data) key(id string) string {
	return path.Join(db.prefix, id)
}

// load parses all available records from the storage.
func (db *data) load() ([]*model.Redirect, error) {
	res := make([]*model.Redirect, 0)
	records, err := db.store.List(db.prefix)

	if err != nil {
		return nil, err
	}

	for _, pair := range records {
		row := &model.Redirect{}

		if err := json.Unmarshal(pair.Value, row); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal record. %s", err)
		}

		res = append(res, row)
	}

	return res, nil
}
