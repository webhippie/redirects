package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/router/middleware/header"
	"github.com/tboerger/redirects/router/middleware/logger"
	"github.com/tboerger/redirects/router/middleware/recovery"
	"github.com/tboerger/redirects/router/middleware/store"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.Use(middleware...)
	e.Use(logger.SetLogger())
	e.Use(recovery.SetRecovery())
	e.Use(store.SetStore())
	e.Use(header.SetVersion())

	e.NoRoute(handler)

	return e
}

func handler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"version": config.Version,
		},
	)
}
