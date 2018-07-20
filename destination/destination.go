package destination

import (
	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Destination is the interface for Destination.
type Destination interface {
	Write(c chan *model.Server) error
	Close() error
	Init() error
}

// LoadConfig will load config to a Destination.
func LoadConfig(d Destination, c config.Endpoint) (err error) {
	// Load options
	content, err := yaml.Marshal(c.Options)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, d)
	if err != nil {
		return
	}

	return
}

// LoadServers will load servers from db to destination.
func LoadServers(d Destination) (err error) {
	c := make(chan *model.Server, 10)

	go model.ListServers(c)

	err = d.Write(c)
	if err != nil {
		return
	}
	return
}
