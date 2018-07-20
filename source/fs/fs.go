package fs

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/Xuanwo/tiresias/model"
)

// Fs will load data from file system.
type Fs struct {
	files []string

	d model.Server

	Path string `yaml:"path"`
}

// Name will return current source's name.
func (f *Fs) Name() string {
	return "fs:" + f.Path
}

// Defaults wil return the default value for server.
func (f *Fs) Defaults() *model.Server {
	return &f.d
}

// Init will initiate Fs.
func (f *Fs) Init() (err error) {
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
func (f *Fs) List(c chan *model.Server) (err error) {
	defer close(c)

	for _, v := range f.files {
		file, err := os.OpenFile(v, os.O_RDONLY, 0600)
		if err != nil {
			log.Printf("Open file failed for %v.", err)
			return err
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		file.Close()

		ts := []model.Server{}
		err = yaml.Unmarshal(content, &ts)
		if err != nil {
			return err
		}

		for k := range ts {
			c <- &ts[k]
		}
	}

	return
}
