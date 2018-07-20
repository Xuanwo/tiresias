package model

import (
	"bytes"

	"github.com/syndtr/goleveldb/leveldb/util"

	"github.com/Xuanwo/tiresias/constants"
	"github.com/Xuanwo/tiresias/contexts"
	"log"
)

// CreateSource will create a source.
func CreateSource(source string) (err error) {
	return contexts.DB.Put(constants.FormatSourceKey(source), nil, nil)
}

// ListSources will list all sources.
func ListSources() (s []string, err error) {
	it := contexts.DB.NewIterator(
		util.BytesPrefix([]byte(constants.KeySourcePrefix)), nil)

	b := it.Seek([]byte(constants.KeySourcePrefix))

	if b {
		key := it.Key()

		if !bytes.HasPrefix(key, []byte(constants.KeySourcePrefix)) {
			b = false
		}

		s = append(s, string(key))

		b = it.Next()
	}

	it.Release()
	err = it.Error()
	return
}

// DeleteSource will delete a source with all it's servers.
func DeleteSource(source string) (err error) {
	c := make(chan *Server, 10)

	go ListServersBySource(source, c)

	for v := range c {
		err = DeleteServer(source, v.Name)
		if err != nil {
			log.Printf("Delete server %s failed for %v.", v.Name, err)
			continue
		}
	}

	return contexts.DB.Delete(constants.FormatSourceKey(source), nil)
}
