package config

// Endpoint is the config for source or destination.
type Endpoint struct {
	Type string `yaml:"type"`

	Options map[string]interface{} `yaml:"options"`
	Default map[string]interface{} `yaml:"default"`
}
