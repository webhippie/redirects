// +build !windows

package cmd

import (
	"github.com/facebookgo/grace/gracehttp"
	"net/http"
)

func startServer(s *http.Server) error {
	return gracehttp.Serve(s)
}
