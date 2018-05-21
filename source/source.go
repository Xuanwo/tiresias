package source

import (
	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Source is the interface for source.
type Source interface {
	Init(c config.Endpoint) (err error)
	List() (s []model.Server, err error)
}
