package handler

import (
	"net/http"

	"github.com/kazukousen/goad/goad/resource"
)

func init() {
	http.Handle("/campaign/", http.HandlerFunc(resource.Chain(resource.XMLHandler(Campaign{}))))
}

// Campaign implements resource.Resource
type Campaign struct {
	resource.Base
}

// Get ...
func (c Campaign) Get(r *http.Request) (resource.Status, interface{}) {
	type Human struct {
		Name string
	}
	boy := &Human{
		Name: "boy",
	}
	return resource.Success(http.StatusOK), boy
}
