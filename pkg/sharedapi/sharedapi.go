package sharedapi

// API is a public API instance.
type API struct{}

// NewAPI returns public API instance.
func NewAPI() *API {
	return &API{}
}
