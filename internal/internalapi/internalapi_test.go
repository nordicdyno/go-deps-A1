package internalapi

import "testing"

func TestNewPrivateAPI(t *testing.T) {
	api := NewPrivateAPI()
	if api == nil {
		t.Error("got nil on NewPrivateAPI call")
	}
}
