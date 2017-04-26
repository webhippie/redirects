package router

import (
	"bytes"
	"net/http"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/webhippie/redirects/model"
	"github.com/webhippie/redirects/router/middleware/source"
	"github.com/webhippie/redirects/store"
)

func handler(c *gin.Context) {
	s := source.Get(c)
	redirects, err := store.GetRedirects(c)

	if err != nil {
		logrus.Errorf("Failed to load redirects. %s", err)

		c.HTML(
			http.StatusInternalServerError,
			"500.html",
			gin.H{},
		)

		return
	}

	sort.Sort(
		model.RedirectsByPriority(redirects),
	)

	logrus.Debugf("Trying to match %s", s.String())

	for _, redirect := range redirects {
		if matched, _ := regexp.MatchString(redirect.Source, s.String()); matched {
			logrus.Debugf("Matched %s with %s", s.String(), redirect.ID)

			tmpl, err := template.New(
				"_",
			).Funcs(
				helpers(),
			).Parse(
				redirect.Destination,
			)

			if err != nil {
				logrus.Errorf("Failed to parse template. %s", err)

				c.HTML(
					http.StatusInternalServerError,
					"500.html",
					gin.H{},
				)

				return
			}

			res := new(bytes.Buffer)

			if err := tmpl.Execute(res, s); err != nil {
				logrus.Errorf("Failed to process template. %s", err)

				c.HTML(
					http.StatusInternalServerError,
					"500.html",
					gin.H{},
				)

				return
			}

			c.Redirect(http.StatusMovedPermanently, res.String())
			return
		}
	}

	c.HTML(
		http.StatusNotFound,
		"404.html",
		gin.H{},
	)
}

func helpers() template.FuncMap {
	return template.FuncMap{
		"split":    strings.Split,
		"join":     strings.Join,
		"toUpper":  strings.ToUpper,
		"toLower":  strings.ToLower,
		"contains": strings.Contains,
		"replace":  strings.Replace,
	}
}
