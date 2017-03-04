package toml

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/tboerger/redirects/model"
	"github.com/tboerger/redirects/store"
	"io/ioutil"
)

// redirectCollection represents the internal storage collection.
type redirectCollection struct {
	Redirects []*model.Redirect `toml:"redirects"`
}

// GetRedirects retrieves all redirects from the TOML store.
func (db *data) GetRedirects() ([]*model.Redirect, error) {
	root, err := loadRedirects(db.dsn)

	if err != nil {
		return nil, err
	}

	return root.Redirects, nil
}

// GetRedirect retrieves a specific redirect from the TOML store.
func (db *data) GetRedirect(id int) (*model.Redirect, error) {
	root, err := loadRedirects(db.dsn)

	if err != nil {
		return nil, err
	}

	if id >= len(root.Redirects) || root.Redirects[id] == nil {
		return nil, store.ErrRedirectNotFound
	}

	return root.Redirects[id], nil
}

// DeleteRedirect deletes a redirect from the TOML store.
func (db *data) DeleteRedirect(id int) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	root, err := loadRedirects(db.dsn)

	if err != nil {
		return err
	}

	if id >= len(root.Redirects) || root.Redirects[id] == nil {
		return store.ErrRedirectNotFound
	}

	root.Redirects = append(
		root.Redirects[:id],
		root.Redirects[id+1:]...,
	)

	return writeRedirects(db.dsn, root)
}

// CreateRedirect creates a redirect on the TOML store.
func (db *data) CreateRedirect(record *model.Redirect) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	root, err := loadRedirects(db.dsn)

	if err != nil {
		return err
	}

	record.ID = len(root.Redirects)

	root.Redirects = append(
		root.Redirects,
		record,
	)

	return writeRedirects(db.dsn, root)
}

// UpdateRedirect updates a redirect on the TOML store.
func (db *data) UpdateRedirect(record *model.Redirect) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	root, err := loadRedirects(db.dsn)

	if err != nil {
		return err
	}

	if record.ID >= len(root.Redirects) || root.Redirects[record.ID] == nil {
		return store.ErrRedirectNotFound
	}

	root.Redirects[record.ID] = record

	return writeRedirects(db.dsn, root)
}

// loadRedirects parses all available records from the storage.
func loadRedirects(dsn string) (*redirectCollection, error) {
	res := &redirectCollection{
		Redirects: make([]*model.Redirect, 0),
	}

	content, err := ioutil.ReadFile(dsn)

	if err != nil {
		return nil, err
	}

	if _, err := toml.Decode(string(content), res); err != nil {
		return nil, err
	}

	for id, record := range res.Redirects {
		record.ID = id
	}

	return res, nil
}

// writeRedirects writes the TOML content back to the storage.
func writeRedirects(dsn string, content *redirectCollection) error {
	for _, record := range content.Redirects {
		record.ID = 0
	}

	buf := new(bytes.Buffer)

	if err := toml.NewEncoder(buf).Encode(content); err != nil {
		return err
	}

	if err := ioutil.WriteFile(dsn, buf.Bytes(), 0640); err != nil {
		return err
	}

	return nil
}
