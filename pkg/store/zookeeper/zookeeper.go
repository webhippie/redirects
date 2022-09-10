package zookeeper

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/kvtools/valkeyrie"
	valkeyrieStore "github.com/kvtools/valkeyrie/store"
	"github.com/kvtools/zookeeper"
	"github.com/webhippie/redirects/pkg/config"
	"github.com/webhippie/redirects/pkg/model"
	"github.com/webhippie/redirects/pkg/store"
)

// data is a basic struct that iplements the Store interface.
type data struct {
	store     valkeyrieStore.Store
	prefix    string
	endpoints []string
}

// Name simply returns the name of the store implementation.
func (db *data) Name() string {
	return "Zookeeper"
}

// Config just returns a simple configuration explanation.
func (db *data) Config() string {
	return fmt.Sprintf("endpoints:%s", strings.Join(db.endpoints, ","))
}

// key generates the new key including the prefix.
func (db *data) key(id string) string {
	return path.Join(db.prefix, "records", id)
}

// load parses all available records from the storage.
func (db *data) load(ctx context.Context) ([]*model.Redirect, error) {
	res := make([]*model.Redirect, 0)
	records, err := db.store.List(ctx, db.prefix, &valkeyrieStore.ReadOptions{})

	if err != nil {
		return nil, err
	}

	for _, pair := range records {
		row := &model.Redirect{}

		if err := json.Unmarshal(pair.Value, row); err != nil {
			return nil, fmt.Errorf("failed to unmarshal record: %w", err)
		}

		res = append(res, row)
	}

	return res, nil
}

// New initializes a new Zookeeper store.
func New(s valkeyrieStore.Store, prefix string, endpoints []string) store.Store {
	return &data{
		store:     s,
		prefix:    prefix,
		endpoints: endpoints,
	}
}

// Load initializes the Zookeeper storage.
func Load(cfg *config.Zookeeper) (store.Store, error) {
	ctx := context.Background()
	prefix := cfg.Prefix

	valkeyrieConfig := &zookeeper.Config{
		ConnectionTimeout: cfg.Timeout * time.Second,
	}

	s, err := valkeyrie.NewStore(
		ctx,
		zookeeper.StoreName,
		cfg.Endpoints,
		valkeyrieConfig,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to init store: %w", err)
	}

	if ok, _ := s.Exists(
		ctx,
		prefix,
		&valkeyrieStore.ReadOptions{},
	); !ok {
		err := s.Put(
			ctx,
			prefix,
			nil,
			&valkeyrieStore.WriteOptions{
				IsDir: true,
			},
		)

		if err != nil {
			return nil, fmt.Errorf("failed to create prefix: %w", err)
		}
	}

	return New(
		s,
		prefix,
		cfg.Endpoints,
	), nil
}
