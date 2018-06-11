package fs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Fs will load data from file system.
type Fs struct {
	files []string

	defaults model.Server

	Path string `yaml:"path"`
}

// Init will initiate Fs.
func (f *Fs) Init(c config.Endpoint) (err error) {
	// Load options
	content, err := yaml.Marshal(c.Options)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, f)
	if err != nil {
		return
	}

	// Load defaults.
	content, err = yaml.Marshal(c.Default)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(content, &f.defaults)
	if err != nil {
		return
	}

	f.files, err = filepath.Glob(f.Path)
	if err != nil {
		return
	}
	if f.files == nil {
		log.Printf("No file matched")
		return
	}

	return
}

// List will list all servers from Fs.
func (f *Fs) List() (s []model.Server, err error) {
	s = []model.Server{}

	for _, v := range f.files {
		file, err := os.OpenFile(v, os.O_RDONLY, 0600)
		if err != nil {
			log.Printf("Open file failed for %v.", err)
			return nil, err
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		ts := []model.Server{}
		err = yaml.Unmarshal(content, &ts)
		if err != nil {
			return nil, err
		}

		s = append(s, ts...)
	}

	for k := range s {
		if f.defaults.User != "" {
			s[k].User = f.defaults.User
		}
		if f.defaults.Port != "" {
			s[k].Port = f.defaults.Port
		}
		if f.defaults.IdentityFile != "" {
			s[k].IdentityFile = f.defaults.IdentityFile
		}
	}
	return
}
