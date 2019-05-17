package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dsielert/redirector/config"
	"github.com/dsielert/redirector/router/middleware/header"
	"github.com/dsielert/redirector/router/middleware/logger"
	"github.com/dsielert/redirector/router/middleware/recovery"
	"github.com/dsielert/redirector/router/middleware/source"
	"github.com/dsielert/redirector/router/middleware/store"
	"github.com/dsielert/redirector/templates"
)

// Load initializes the routing of the application.
func Load(middleware ...gin.HandlerFunc) http.Handler {
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e := gin.New()

	e.SetHTMLTemplate(
		templates.Load(),
	)

	e.Use(middleware...)
	e.Use(logger.SetLogger())
	e.Use(recovery.SetRecovery())
	e.Use(store.SetStore())
	e.Use(header.SetVersion())
	e.Use(source.SetSource())

	e.NoRoute(handler)

	return e
}
