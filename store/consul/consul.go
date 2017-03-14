package consul

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/docker/libkv"
	libkvStore "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"github.com/jackspirou/syscerts"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/model"
	"github.com/tboerger/redirects/store"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

// init simply registers Consul on the libkv library.
func init() {
	consul.Register()
}

// data is a basic struct that iplements the Store interface.
type data struct {
	store     libkvStore.Store
	prefix    string
	endpoints []string
}

// Name simply returns the name of the store implementation.
func (db *data) Name() string {
	return "Consul"
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

// New initializes a new Consul store.
func New(s libkvStore.Store, prefix string, endpoints []string) store.Store {
	return &data{
		store:     s,
		prefix:    prefix,
		endpoints: endpoints,
	}
}

// Load initializes the Consul storage.
func Load() (store.Store, error) {
	prefix := config.Consul.Prefix

	cfg := &libkvStore.Config{
		ConnectionTimeout: config.Consul.Timeout * time.Second,
	}

	if config.Consul.Cert != "" && config.Consul.Key != "" {
		pool, err := pool()

		if err != nil {
			return nil, fmt.Errorf("Failed to init cert pool. %s", err)
		}

		cert, err := cert()

		if err != nil {
			return nil, fmt.Errorf("Failed to init SSL cert. %s", err)
		}

		key, err := key()

		if err != nil {
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to init SSL key. %s", err))
		}

		keypair, err := tls.X509KeyPair(cert, key)

		if err != nil {
			return nil, fmt.Errorf("Failed to parse keypair. %s", err)
		}

		cfg.TLS = &tls.Config{
			Certificates:       []tls.Certificate{keypair},
			RootCAs:            pool,
			InsecureSkipVerify: config.Consul.SkipVerify,
		}
	}

	s, err := libkv.NewStore(
		libkvStore.CONSUL,
		config.Consul.Endpoints,
		cfg,
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to init store. %s", err)
	}

	if ok, _ := s.Exists(prefix); !ok {
		err := s.Put(
			prefix,
			nil,
			&libkvStore.WriteOptions{
				IsDir: true,
			},
		)

		if err != nil {
			return nil, fmt.Errorf("Failed to create prefix. %s", err)
		}
	}

	return New(
		s,
		prefix,
		config.Consul.Endpoints,
	), nil
}

// pool initializes the CA cert pool from system and custom CA file or flag.
func pool() (*x509.CertPool, error) {
	pool := syscerts.SystemRootsPool()

	if config.Consul.CA != "" {
		var ca []byte

		if _, err := os.Stat(config.Consul.CA); err == nil {
			ca, err = ioutil.ReadFile(config.Consul.CA)

			if err != nil {
				return nil, fmt.Errorf("Failed to read CA certificate. %s", err)
			}
		} else {
			ca = []byte(config.Consul.CA)
		}

		pool.AppendCertsFromPEM(ca)
	}

	return pool, nil
}

// key loads the SSL key from file or flag.
func key() ([]byte, error) {
	if _, err := os.Stat(config.Consul.Key); err == nil {
		return ioutil.ReadFile(config.Consul.Key)
	}

	return []byte(config.Consul.Key), nil
}

// cert loads the SSL certificate from file or flag.
func cert() ([]byte, error) {
	if _, err := os.Stat(config.Consul.Cert); err == nil {
		return ioutil.ReadFile(config.Consul.Cert)
	}

	return []byte(config.Consul.Cert), nil
}
