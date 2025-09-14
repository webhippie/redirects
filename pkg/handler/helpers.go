package handler

import (
	"html/template"
	"strings"
)

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
	return strings.ReplaceAll(s3, s1, s2)
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
