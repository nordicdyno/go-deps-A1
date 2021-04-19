package sharedapi

import "testing"

func TestNewAPI(t *testing.T) {
	api := NewAPI()
	if api == nil {
		t.Error("got nil on NewAPI call")
	}
}
