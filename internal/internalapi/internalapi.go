package internalapi

import "github.com/nordicdyno/go-deps-A1/pkg/sharedapi"

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
