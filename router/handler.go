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
	"github.com/dsielert/redirector/model"
	"github.com/dsielert/redirector/router/middleware/source"
	"github.com/dsielert/redirector/store"
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

			logrus.Debugf("Redirecting %s to %s", s.String(), res.String())
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
		"toUpper":     strings.ToUpper,
		"toLower":     strings.ToLower,
		"contains":    contains,
		"replace":     replace,
		"split":       split,
		"join":        join,
		"firstString": firstString,
		"lastString":  lastString,
	}
}

func replace(s1, s2, s3 string) string {
	return strings.Replace(s3, s1, s2, -1)
}

func contains(s1, s2 string) bool {
	return strings.Contains(s2, s1)
}

func split(s1, s2 string) []string {
	return strings.Split(s2, s1)
}

func join(s1 string, s2 []string) string {
	return strings.Join(s2, s1)
}

func firstString(s1 []string) string {
	if len(s1) < 1 {
		return ""
	}

	return s1[0]
}

func lastString(s1 []string) string {
	if len(s1) < 1 {
		return ""
	}

	return s1[len(s1)-1]
}
