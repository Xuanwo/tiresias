package model

import (
	"bytes"
	"log"

	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/vmihailenco/msgpack"

	"github.com/Xuanwo/tiresias/constants"
	"github.com/Xuanwo/tiresias/contexts"
)

// Server stores all information that user needed to access a server.
type Server struct {
	Name         string `yaml:"name" msgpack:"n"`
	Address      string `yaml:"address" msgpack:"a"`
	Port         string `yaml:"port" msgpack:"p"`
	User         string `yaml:"user" msgpack:"u"`
	IdentityFile string `yaml:"identity_file" msgpack:"if"`
}

// CreateServer will create a server.
func CreateServer(source string, s *Server) (err error) {
	content, err := msgpack.Marshal(s)
	if err != nil {
		log.Printf("Msgpack marshal failed for %v.", err)
		return
	}

	return contexts.DB.Put(constants.FormatServerKey(source, s.Name), content, nil)
}

// ListServers will list all servers.
func ListServers(c chan *Server) (err error) {
	defer close(c)

	it := contexts.DB.NewIterator(
		util.BytesPrefix([]byte(constants.KeyServerPrefix)), nil)

	b := it.Seek([]byte(constants.KeyServerPrefix))

	for b {
		key := it.Key()

		if !bytes.HasPrefix(key, []byte(constants.KeyServerPrefix)) {
			b = false
		}

		s := &Server{}
		err = msgpack.Unmarshal(it.Value(), s)
		if err != nil {
			return
		}

		c <- s

		b = it.Next()
	}

	it.Release()
	err = it.Error()
	if err != nil {
		log.Fatalf("List servers failed for %v.", err)
		return
	}
	return
}
