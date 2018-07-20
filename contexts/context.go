package contexts

import (
	"github.com/Xuanwo/tiresias/config"
	"github.com/Xuanwo/tiresias/db"
)

var (
	// DB holds the database connection.
	DB *db.Database
)

// SetupContexts will set contexts.
func SetupContexts(c *config.Config) (err error) {
	// Setup db.
	DB, err = db.NewDB(&db.DatabaseOptions{
		Address: c.Database,
	})
	if err != nil {
		return
	}
	return
}
