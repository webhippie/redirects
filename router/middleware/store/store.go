package store

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/store"
	"github.com/tboerger/redirects/store/json"
	"github.com/tboerger/redirects/store/yaml"
)

// SetStore injects the storage into the context.
func SetStore() gin.HandlerFunc {
	var (
		db store.Store
	)

	switch config.Storage.Driver {
	case "yaml":
		db = yaml.Load()
	case "json":
		db = json.Load()
	}

	logrus.Infof("Using storage driver %s", config.Storage.Driver)
	logrus.Infof("Using storage DSN %s", config.Storage.DSN)

	return func(c *gin.Context) {
		store.ToContext(c, db)
		c.Next()
	}
}
