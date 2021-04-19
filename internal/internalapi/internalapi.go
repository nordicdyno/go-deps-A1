package internalapi

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/nordicdyno/go-deps-A1/pkg/sharedapi"
	godepsprivate "github.com/nordicdyno/go-deps-private"
)

// PrivateAPI is a private API instance.
type PrivateAPI struct {
	api *sharedapi.API
}

// NewPrivateAPI returns internal API instance.
func NewPrivateAPI() *PrivateAPI {
	return &PrivateAPI{
		api: sharedapi.NewAPI(),
	}
}

// LiteConnect creates *gorm.DB instance binded to SQLite file.
//
// NOTE: You can also use file::memory:?cache=shared instead of a path to a file.
// This will tell SQLite to use a temporary database in system memory.
func (p *PrivateAPI) LiteConnect(dbfile string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dbfile), &gorm.Config{})
}

func topSecret() {
	godepsprivate.TopSecret()
}
