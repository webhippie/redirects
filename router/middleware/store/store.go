package store

import (
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
	"strings"
)

// SetStore injects the storage into the context.
func SetStore() gin.HandlerFunc {
	var (
		db store.Store
	)

	switch {
	case config.YAML.Enabled:
		logrus.Infof("Using storage driver: YAML")
		logrus.Infof("Using storage file: %s", config.YAML.File)
		db = yaml.Load()
	case config.JSON.Enabled:
		logrus.Infof("Using storage driver: JSON")
		logrus.Infof("Using storage file: %s", config.JSON.File)
		db = json.Load()
	case config.TOML.Enabled:
		logrus.Infof("Using storage driver: TOML")
		logrus.Infof("Using storage file: %s", config.TOML.File)
		db = toml.Load()
	case config.Zookeeper.Enabled:
		logrus.Infof("Using storage driver: Zookeeper")
		logrus.Infof("Using storage endpoints: %s", strings.Join(config.Zookeeper.Endpoints, ", "))
		db = zookeeper.Load()
	case config.Etcd.Enabled:
		logrus.Infof("Using storage driver: Etcd")
		logrus.Infof("Using storage endpoints: %s", strings.Join(config.Etcd.Endpoints, ", "))
		db = etcd.Load()
	case config.Consul.Enabled:
		logrus.Infof("Using storage driver: Consul")
		logrus.Infof("Using storage endpoints: %s", strings.Join(config.Consul.Endpoints, ", "))
		db = consul.Load()
	}

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
