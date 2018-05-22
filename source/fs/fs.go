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
}

// Init will initiate Fs.
func (f *Fs) Init(c config.Endpoint) (err error) {
	f.files, err = filepath.Glob(c.Path)
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
	return
}
