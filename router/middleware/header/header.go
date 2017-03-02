package header

import (
	"github.com/gin-gonic/gin"
	"github.com/tboerger/redirects/config"
)

// SetVersion writes the current API version to the headers.
func SetVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-REDIRECTS-VERSION", config.Version)
		c.Next()
	}
}
