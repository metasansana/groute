package groute

import (
	"net/http"
)

//StaticRoute provides static routing support.
type StaticRoute struct {
	path   string
	action Callback
}

//Trigger implements the method from the Route interface.
func (self *StaticRoute) Trigger(route string, w http.ResponseWriter, r *http.Request) bool {

	if self.path == route {

		self.action(w, r)

		return true

	}

	return false

}

//NewStaticRoute constructor.
func NewStaticRoute(path string, callback Callback) Route {

	return &StaticRoute{path, callback}

}
