package etcd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/jackspirou/syscerts"
	etcd "github.com/kvtools/etcdv3"
	"github.com/kvtools/valkeyrie"
	valkeyrieStore "github.com/kvtools/valkeyrie/store"
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
	return "Etcd"
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

// New initializes a new Etcd store.
func New(s valkeyrieStore.Store, prefix string, endpoints []string) store.Store {
	return &data{
		store:     s,
		prefix:    prefix,
		endpoints: endpoints,
	}
}

// Load initializes the Etcd storage.
func Load(cfg *config.Etcd) (store.Store, error) {
	ctx := context.Background()
	prefix := cfg.Prefix

	valkeyrieConfig := &etcd.Config{
		ConnectionTimeout: cfg.Timeout * time.Second,
		Username:          cfg.Username,
		Password:          cfg.Password,
	}

	if cfg.Cert != "" && cfg.Key != "" {
		pool, err := pool(cfg)

		if err != nil {
			return nil, fmt.Errorf("failed to init cert pool: %w", err)
		}

		cert, err := cert(cfg)

		if err != nil {
			return nil, fmt.Errorf("failed to init SSL cert: %w", err)
		}

		key, err := key(cfg)

		if err != nil {
			return nil, fmt.Errorf("failed to init SSL key: %w", err)
		}

		keypair, err := tls.X509KeyPair(cert, key)

		if err != nil {
			return nil, fmt.Errorf("failed to parse keypair: %w", err)
		}

		valkeyrieConfig.TLS = &tls.Config{
			Certificates:       []tls.Certificate{keypair},
			RootCAs:            pool,
			InsecureSkipVerify: cfg.SkipVerify,
		}
	}

	s, err := valkeyrie.NewStore(
		ctx,
		etcd.StoreName,
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

// pool initializes the CA cert pool from system and custom CA file or flag.
func pool(cfg *config.Etcd) (*x509.CertPool, error) {
	pool := syscerts.SystemRootsPool()

	if cfg.CA != "" {
		var ca []byte

		if _, err := os.Stat(cfg.CA); err == nil {
			ca, err = os.ReadFile(cfg.CA)

			if err != nil {
				return nil, fmt.Errorf("failed to read CA certificate: %w", err)
			}
		} else {
			ca = []byte(cfg.CA)
		}

		pool.AppendCertsFromPEM(ca)
	}

	return pool, nil
}

// key loads the SSL key from file or flag.
func key(cfg *config.Etcd) ([]byte, error) {
	if _, err := os.Stat(cfg.Key); err == nil {
		return os.ReadFile(cfg.Key)
	}

	return []byte(cfg.Key), nil
}

// cert loads the SSL certificate from file or flag.
func cert(cfg *config.Etcd) ([]byte, error) {
	if _, err := os.Stat(cfg.Cert); err == nil {
		return os.ReadFile(cfg.Cert)
	}

	return []byte(cfg.Cert), nil
}
