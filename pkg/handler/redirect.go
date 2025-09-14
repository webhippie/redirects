package handler

import (
	"bytes"
	"html/template"
	"net/http"
	"regexp"
	"sort"

	"github.com/rs/zerolog/log"
	"github.com/webhippie/redirects/pkg/config"
	"github.com/webhippie/redirects/pkg/middleware/source"
	"github.com/webhippie/redirects/pkg/model"
	"github.com/webhippie/redirects/pkg/store"
	"github.com/webhippie/redirects/pkg/templates"
)

// Redirect is used to handle all the redirects.
func Redirect(cfg *config.Config, storage store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		redirects, err := storage.GetRedirects(req.Context())

		if err != nil {
			log.Error().
				Err(err).
				Msg("Failed to load redirects")

			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")

			if err := templates.Load(cfg).ExecuteTemplate(
				w,
				"500.tmpl",
				nil,
			); err != nil {
				http.Error(
					w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}

			return
		}

		sort.Sort(
			model.RedirectsByPriority(redirects),
		)

		s := source.Get(req.Context())

		log.Debug().
			Str("source", s.String()).
			Msg("Trying to match URL")

		for _, redirect := range redirects {
			if matched, _ := regexp.MatchString(redirect.Source, s.String()); matched {
				log.Debug().
					Str("source", s.String()).
					Str("id", redirect.ID).
					Msg("Found a match for URL")

				tmpl, err := template.New(
					"_",
				).Funcs(
					helpers(),
				).Parse(
					redirect.Destination,
				)

				if err != nil {
					log.Error().
						Err(err).
						Str("source", s.String()).
						Str("id", redirect.ID).
						Err(err).
						Msg("Failed to parse template")

					w.WriteHeader(http.StatusInternalServerError)
					w.Header().Set("Content-Type", "text/html; charset=utf-8")

					if err := templates.Load(cfg).ExecuteTemplate(
						w,
						"500.tmpl",
						nil,
					); err != nil {
						http.Error(
							w,
							http.StatusText(http.StatusInternalServerError),
							http.StatusInternalServerError,
						)
					}

					return
				}

				res := new(bytes.Buffer)

				if err := tmpl.Execute(res, s); err != nil {
					log.Error().
						Err(err).
						Str("source", s.String()).
						Str("id", redirect.ID).
						Err(err).
						Msg("Failed to process template")

					w.WriteHeader(http.StatusInternalServerError)
					w.Header().Set("Content-Type", "text/html; charset=utf-8")

					if err := templates.Load(cfg).ExecuteTemplate(
						w,
						"500.tmpl",
						nil,
					); err != nil {
						http.Error(
							w,
							http.StatusText(http.StatusInternalServerError),
							http.StatusInternalServerError,
						)
					}

					return
				}

				log.Debug().
					Err(err).
					Str("source", s.String()).
					Str("id", redirect.ID).
					Str("target", res.String()).
					Msg("Redirecting to matched target")

				http.Redirect(
					w,
					req,
					res.String(),
					http.StatusMovedPermanently,
				)

				return
			}
		}

		log.Debug().
			Str("source", s.String()).
			Msg("Failed to find a match")

		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if err := templates.Load(cfg).ExecuteTemplate(
			w,
			"404.tmpl",
			nil,
		); err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusNotFound),
				http.StatusNotFound,
			)
		}
	}
}
