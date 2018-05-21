package config

// Endpoint is the config for source or destination.
type Endpoint struct {
	Type    string            `yaml:"type"`
	Path    string            `yaml:"path"`
	Options map[string]string `yaml:"options"`
}
