package templates

import (
	"embed"
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/webhippie/redirects/pkg/config"
)

var (
	//go:embed dist/*
	embeddedTemplates embed.FS

	allowedExtensions = []string{
		".html",
		".tmpl",
	}
)

// Load initializes the template files.
func Load(cfg *config.Config) *template.Template {
	tpls := template.New("")

	err := fs.WalkDir(embeddedTemplates, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if forbiddenExtension(filepath.Ext(d.Name())) {
			return nil
		}

		content, err := fs.ReadFile(
			embeddedTemplates,
			p,
		)

		if err != nil {
			return err
		}

		tpls.New(
			strings.TrimPrefix(
				d.Name(),
				"dist/",
			),
		).Parse(
			string(content),
		)

		return nil
	})

	if err != nil {
		log.Warn().
			Err(err).
			Msg("Failed to parse builtin templates")
	}

	if cfg.Server.Templates != "" {
		if stat, err := os.Stat(cfg.Server.Templates); os.IsNotExist(err) || !stat.IsDir() {
			log.Warn().
				Err(err).
				Msg("Custom templates directory does not exit")

			return tpls
		}

		err := filepath.Walk(cfg.Server.Templates, func(p string, d fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			if forbiddenExtension(filepath.Ext(d.Name())) {
				return nil
			}

			content, err := ioutil.ReadFile(
				p,
			)

			if err != nil {
				return err
			}

			tpls.New(
				strings.TrimPrefix(
					strings.TrimPrefix(
						d.Name(),
						cfg.Server.Templates,
					),
					"/",
				),
			).Parse(
				string(content),
			)

			return nil
		})

		if err != nil {
			log.Warn().
				Err(err).
				Msg("Failed to parse custom templates")
		}
	}

	return tpls
}

func forbiddenExtension(ext string) bool {
	for _, allowed := range allowedExtensions {
		if ext == allowed {
			return false
		}
	}

	return true
}
