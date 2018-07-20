package source

import (
	"log"

	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Source is the interface for source.
type Source interface {
	Defaults() *model.Server
	List(chan *model.Server) (err error)
	Name() string
	Init() error
}

// LoadConfig will load config to a source.
func LoadConfig(s Source, c config.Endpoint) (err error) {
	// Load options
	content, err := yaml.Marshal(c.Options)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, s)
	if err != nil {
		return
	}

	// Load defaults.
	content, err = yaml.Marshal(c.Default)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, s.Defaults())
	if err != nil {
		return
	}

	return
}

// SaveServers will save servers to db.
func SaveServers(s Source) (err error) {
	c := make(chan *model.Server, 10)
	d := s.Defaults()

	// Create source before list it.
	err = model.CreateSource(s.Name())
	if err != nil {
		log.Printf("Save source to db failed for %v.", err)
		return
	}

	go s.List(c)

	for v := range c {
		if d.User != "" {
			v.User = d.User
		}
		if d.Port != "" {
			v.Port = d.Port
		}
		if d.IdentityFile != "" {
			v.IdentityFile = d.IdentityFile
		}

		// We don't need to break the loop while creating server failed here.
		err = model.CreateServer(s.Name(), v)
		if err != nil {
			log.Printf("Save server to db failed for %v.", err)
		}
	}

	return
}
