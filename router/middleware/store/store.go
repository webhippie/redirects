package store

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/webhippie/redirects/config"
	"github.com/webhippie/redirects/store"
	"github.com/webhippie/redirects/store/consul"
	"github.com/webhippie/redirects/store/etcd"
	"github.com/webhippie/redirects/store/json"
	"github.com/webhippie/redirects/store/toml"
	"github.com/webhippie/redirects/store/yaml"
	"github.com/webhippie/redirects/store/zookeeper"
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
		if err != nil {
			c.HTML(
				http.StatusInternalServerError,
				"500.html",
				gin.H{},
			)

			c.Abort()
		} else {
			store.ToContext(c, db)
			c.Next()
		}
	}
}
