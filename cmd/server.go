package cmd

import (
	"crypto/tls"
	"github.com/Sirupsen/logrus"
	"github.com/tboerger/redirects/config"
	"github.com/tboerger/redirects/router"
	"golang.org/x/crypto/acme/autocert"
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"time"

	// Import pprof for optional debugging
	_ "net/http/pprof"
)

// Server provides the sub-command to start the server.
func Server() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start the integrated server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "addr",
				Value:       "0.0.0.0:8080",
				Usage:       "Address to bind the server",
				EnvVars:     []string{"KLEISTER_SERVER_ADDR"},
				Destination: &config.Server.Addr,
			},
			&cli.BoolFlag{
				Name:        "pprof",
				Usage:       "Enable pprof debugging server",
				EnvVars:     []string{"KLEISTER_SERVER_PPROF"},
				Destination: &config.Server.Pprof,
			},
			&cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVars:     []string{"KLEISTER_SERVER_CERT"},
				Destination: &config.Server.Cert,
			},
			&cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVars:     []string{"KLEISTER_SERVER_KEY"},
				Destination: &config.Server.Key,
			},
			&cli.StringFlag{
				Name:        "templates",
				Value:       "",
				Usage:       "Path to custom templates",
				EnvVars:     []string{"KLEISTER_SERVER_TEMPLATES"},
				Destination: &config.Server.Templates,
			},
			&cli.BoolFlag{
				Name:        "letsencrypt",
				Usage:       "Enable Let's Encrypt SSL",
				EnvVars:     []string{"KLEISTER_SERVER_LETSENCRYPT"},
				Destination: &config.Server.LetsEncrypt,
			},
		},
		Action: func(c *cli.Context) error {
			logrus.Infof("Starting the server on %s", config.Server.Addr)

			var (
				server *http.Server
			)

			if config.Server.Pprof {
				logrus.Infof("Starting the debugger on localhost:6060")

				go func() {
					if err := http.ListenAndServe("localhost:6060", nil); err != nil {
						logrus.Info(err)
					}
				}()
			}

			if config.Server.LetsEncrypt || (config.Server.Cert != "" && config.Server.Key != "") {
				curves := []tls.CurveID{
					tls.CurveP521,
					tls.CurveP384,
					tls.CurveP256,
				}

				ciphers := []uint16{
					tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				}

				cfg := &tls.Config{
					PreferServerCipherSuites: true,
					MinVersion:               tls.VersionTLS12,
					CurvePreferences:         curves,
					CipherSuites:             ciphers,
				}

				if config.Server.LetsEncrypt {
					certManager := autocert.Manager{
						Prompt: autocert.AcceptTOS,
						Cache:  autocert.DirCache("certs"), // TODO: Implement the interface matching the storage type
					}

					cfg.GetCertificate = certManager.GetCertificate
				} else {
					cert, err := tls.LoadX509KeyPair(
						config.Server.Cert,
						config.Server.Key,
					)

					if err != nil {
						logrus.Fatal("Failed to load SSL certificates. %s", err)
					}

					cfg.Certificates = []tls.Certificate{
						cert,
					}
				}

				server = &http.Server{
					Addr:         config.Server.Addr,
					Handler:      router.Load(),
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
					TLSConfig:    cfg,
				}
			} else {
				server = &http.Server{
					Addr:         config.Server.Addr,
					Handler:      router.Load(),
					ReadTimeout:  5 * time.Second,
					WriteTimeout: 10 * time.Second,
				}
			}

			if err := startServer(server); err != nil {
				logrus.Fatal(err)
			}

			return nil
		},
	}
}
