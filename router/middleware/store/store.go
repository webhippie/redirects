package store

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"github.com/tboerger/redirects/store/consul"
	"github.com/tboerger/redirects/store/etcd"
	"github.com/tboerger/redirects/store/json"
	"github.com/tboerger/redirects/store/toml"
	"github.com/tboerger/redirects/store/yaml"
	"github.com/tboerger/redirects/store/zookeeper"
)

// SetStore injects the storage into the context.
func SetStore() gin.HandlerFunc {
	var (
		err error
		db  store.Store
	)

	switch {
	case config.YAML.Enabled:
		db, err = yaml.Load()
	case config.JSON.Enabled:
		db, err = json.Load()
	case config.TOML.Enabled:
		db, err = toml.Load()
	case config.Zookeeper.Enabled:
		db, err = zookeeper.Load()
	case config.Etcd.Enabled:
		db, err = etcd.Load()
	case config.Consul.Enabled:
		db, err = consul.Load()
	default:
		err = fmt.Errorf("No storage method specified")
	}

	if db != nil {
		logrus.Infof("Using storage driver: %s", db.Name())
		logrus.Infof("Using storage config: %s", db.Config())
	} else {
		logrus.Errorf("%s", err)
	}

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
