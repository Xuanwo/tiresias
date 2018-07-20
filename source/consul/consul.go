package consul

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Xuanwo/tiresias/model"
	"github.com/hashicorp/consul/api"
)

// Consul will load data from consul.
type Consul struct {
	client  *api.Client
	catalog *api.Catalog

	d model.Server

	Address    string `yaml:"address"`
	Schema     string `yaml:"schema"`
	Datacenter string `yaml:"datacenter"`
	Prefix     string `yaml:"prefix"`
}

// Name will return current source's name.
func (c *Consul) Name() string {
	return fmt.Sprintf("consul:%s:%s", c.Address, c.Datacenter)
}

// Defaults will return the default value for server.
func (c *Consul) Defaults() *model.Server {
	return &c.d
}

// Init will initiate Fs.
func (c *Consul) Init() (err error) {
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

// List will list all servers from Consul.
func (c *Consul) List(ch chan *model.Server) (err error) {
	defer close(ch)

	nodes, _, err := c.catalog.Nodes(nil)
	if err != nil {
		return
	}
	for _, v := range nodes {
		ch <- &model.Server{
			Name:    c.Prefix + v.Node,
			Address: v.Address,
		}
	}
	return
}
