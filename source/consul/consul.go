package consul

import (
	"log"
	"net"
	"time"

	"github.com/hashicorp/consul/api"
	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Consul will load data from consul.
type Consul struct {
	client  *api.Client
	catalog *api.Catalog

	defaults model.Server

	Address    string `yaml:"address"`
	Schema     string `yaml:"schema"`
	Datacenter string `yaml:"datacenter"`
	Prefix     string `yaml:"prefix"`
}

// Init will initiate Fs.
func (c *Consul) Init(e config.Endpoint) (err error) {
	// Load options
	content, err := yaml.Marshal(e.Options)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, c)
	if err != nil {
		return
	}

	// Load defaults.
	content, err = yaml.Marshal(e.Default)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &c.defaults)
	if err != nil {
		return
	}

	// Check if consul reachable before connect.
	conn, err := net.DialTimeout("tcp", c.Address, 3*time.Second)
	if err != nil {
		log.Printf("Consul %s connect failed for %v.", c.Address, err)
		return
	}
	conn.Close()

	cc := api.DefaultConfig()
	cc.Address = c.Address
	cc.Datacenter = c.Datacenter
	cc.Scheme = c.Schema
	c.client, err = api.NewClient(cc)
	if err != nil {
		return
	}

	c.catalog = c.client.Catalog()
	return
}

// List will list all servers from Fs.
func (c *Consul) List() (s []model.Server, err error) {
	nodes, _, err := c.catalog.Nodes(nil)
	if err != nil {
		return
	}
	s = make([]model.Server, len(nodes))
	for k, v := range nodes {
		s[k] = model.Server{
			Name:    c.Prefix + v.Node,
			Address: v.Address,
		}
	}

	for k := range s {
		if c.defaults.User != "" {
			s[k].User = c.defaults.User
		}
		if c.defaults.Port != "" {
			s[k].Port = c.defaults.Port
		}
		if c.defaults.IdentityFile != "" {
			s[k].IdentityFile = c.defaults.IdentityFile
		}
	}

	return
}
