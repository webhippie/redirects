package command

import (
	"strings"
	"text/template"
)

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{
	"split":    strings.Split,
	"join":     strings.Join,
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"contains": strings.Contains,
	"replace":  strings.Replace,
}
