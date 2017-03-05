package etcd

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/docker/libkv"
	libkvStore "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/etcd"
	"github.com/jackspirou/syscerts"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"io/ioutil"
	"os"
	"time"
)

// init simply registers Etcd on the libkv library.
func init() {
	etcd.Register()
}

// data is a basic struct that iplements the Store interface.
type data struct {
	store  libkvStore.Store
	prefix string
}

// New initializes a new Etcd store.
func New(s libkvStore.Store, prefix string) store.Store {
	return &data{
		store:  s,
		prefix: prefix,
	}
}

// Load initializes the Etcd storage.
func Load() store.Store {
	prefix := config.Etcd.Prefix

	cfg := &libkvStore.Config{
		ConnectionTimeout: config.Etcd.Timeout * time.Second,
		Username:          config.Etcd.Username,
		Password:          config.Etcd.Password,
	}

	if config.Etcd.Cert != "" && config.Etcd.Key != "" {
		pool, err := pool()

		if err != nil {
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to init cert pool. %s", err))
		}

		cert, err := cert()

		if err != nil {
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to init SSL cert. %s", err))
		}

		key, err := key()

		if err != nil {
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to init SSL key. %s", err))
		}

		keypair, err := tls.X509KeyPair(cert, key)

		if err != nil {
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to parse keypair. %s", err))
		}

		cfg.TLS = &tls.Config{
			Certificates:       []tls.Certificate{keypair},
			RootCAs:            pool,
			InsecureSkipVerify: config.Etcd.SkipVerify,
		}
	}

	s, err := libkv.NewStore(
		libkvStore.ETCD,
		config.Etcd.Endpoints,
		cfg,
	)

	if err != nil {
		// TODO: Handle this error properly
		panic(fmt.Sprintf("TODO: Failed to init store. %s", err))
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
			// TODO: Handle this error properly
			panic(fmt.Sprintf("TODO: Failed to create prefix. %s", err))
		}
	}

	return New(
		s,
		prefix,
	)
}

// pool initializes the CA cert pool from system and custom CA file or flag.
func pool() (*x509.CertPool, error) {
	pool := syscerts.SystemRootsPool()

	if config.Etcd.CA != "" {
		var ca []byte

		if _, err := os.Stat(config.Etcd.CA); err == nil {
			ca, err = ioutil.ReadFile(config.Etcd.CA)

			if err != nil {
				return nil, fmt.Errorf("Failed to read CA certificate. %s", err)
			}
		} else {
			ca = []byte(config.Etcd.CA)
		}

		pool.AppendCertsFromPEM(ca)
	}

	return pool, nil
}

// key loads the SSL key from file or flag.
func key() ([]byte, error) {
	if _, err := os.Stat(config.Etcd.Key); err == nil {
		return ioutil.ReadFile(config.Etcd.Key)
	}

	return []byte(config.Etcd.Key), nil
}

// cert loads the SSL certificate from file or flag.
func cert() ([]byte, error) {
	if _, err := os.Stat(config.Etcd.Cert); err == nil {
		return ioutil.ReadFile(config.Etcd.Cert)
	}

	return []byte(config.Etcd.Cert), nil
}
