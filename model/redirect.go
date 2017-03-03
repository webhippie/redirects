package model

// Redirect represents a redirect model definition.
type Redirect struct {
	ID          int    `yaml:"id,omitempty" json:"id,omitempty"`
	Source      string `yaml:"source" json:"source"`
	Destination string `yaml:"destination" json:"destination"`
}
