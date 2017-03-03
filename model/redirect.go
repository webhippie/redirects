package model

import (
	"strings"
)

// Redirect represents a redirect model definition.
type Redirect struct {
	ID          int    `yaml:"id,omitempty" json:"id,omitempty"`
	Source      string `yaml:"source" json:"source"`
	Destination string `yaml:"destination" json:"destination"`
}

// Contains checks if a needle is part of the redirect.
func (r *Redirect) Contains(needle string) bool {
	if strings.Contains(r.Source, needle) {
		return true
	}

	if strings.Contains(r.Destination, needle) {
		return true
	}

	return false
}
