package destination

import (
	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/model"
)

// Destnation is the interface for destnation.
type Destnation interface {
	Init(c config.Endpoint) (err error)
	Write(s ...model.Server) (n int, err error)
}
