package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Config is the config for tiresias.
type Config struct {
	Hosts     string `yaml:"hosts"`
	SSHConfig string `yaml:"ssh_config"`

	Servers []Server `yaml:"servers"`
}

// New will create a new config object.
func New() (*Config, error) {
	return &Config{}, nil
}

// LoadFromFilePath will load config from file.
func (c *Config) LoadFromFilePath(filePath string) (err error) {

	f, err := os.OpenFile(filePath, os.O_RDONLY, 0600)
	if err != nil {
		log.Printf("Open file failed for %v.", err)
		return
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	return c.LoadFromContent(content)
}

// LoadFromContent will load config from content.
func (c *Config) LoadFromContent(content []byte) error {
	return yaml.Unmarshal(content, c)
}
