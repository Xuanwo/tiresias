package fs

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Fs will load data from file system.
type Fs struct {
	Path string
}

// Init will initiate Fs.
func (f *Fs) Init(c config.Endpoint) (err error) {
	f.Path = c.Path
	return
}

// List will list all servers from Fs.
func (f *Fs) List() (s []model.Server, err error) {
	file, err := os.OpenFile(f.Path, os.O_RDONLY, 0600)
	if err != nil {
		log.Printf("Open file failed for %v.", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &s)
	if err != nil {
		return
	}
	return
}
