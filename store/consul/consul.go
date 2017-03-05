package consul

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/docker/libkv"
	libkvStore "github.com/docker/libkv/store"
	"github.com/docker/libkv/store/consul"
	"github.com/jackspirou/syscerts"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"io/ioutil"
	"os"
	"time"
)

// init simply registers Consul on the libkv library.
func init() {
	consul.Register()
}

// data is a basic struct that iplements the Store interface.
type data struct {
	store  libkvStore.Store
	prefix string
}

// New initializes a new Consul store.
func New(s libkvStore.Store, prefix string) store.Store {
	return &data{
		store:  s,
		prefix: prefix,
	}
}

// Load initializes the Consul storage.
func Load() store.Store {
	prefix := config.Consul.Prefix

	cfg := &libkvStore.Config{
		ConnectionTimeout: config.Consul.Timeout * time.Second,
	}

	if config.Consul.Cert != "" && config.Consul.Key != "" {
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
			InsecureSkipVerify: config.Consul.SkipVerify,
		}
	}

	s, err := libkv.NewStore(
		libkvStore.CONSUL,
		config.Consul.Endpoints,
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
