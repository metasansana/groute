package groute

import (
	"net/http"
	"regexp"
)

//Callback coresponds to Callback's signature
type Callback func(w http.ResponseWriter, r *Request)

//Route is an inteface used to setup routing.
type Route interface {
	Trigger(route string, w http.ResponseWriter, r *http.Request) bool
}

//RegexRoute provides regex routing support.
type RegexRoute struct {
	path   *regexp.Regexp
	action Callback
}

//Trigger implements the method from the Route interface.
func (self *RegexRoute) Trigger(route string, w http.ResponseWriter, r *http.Request) bool {

	if params := self.path.FindStringSubmatch(route); len(params) > 0 {

		self.action(w, newRequest(r, params))

		return true

	}
	return false
}

//StaticRoute provides static routing support.
type StaticRoute struct {
	path   string
	action Callback
}

//Trigger implements the method from the Route interface.
func (self *StaticRoute) Trigger(route string, w http.ResponseWriter, r *http.Request) bool {

	if self.path == route {

		self.action(w, newRequest(r, make([]string, 0)))

		return true

	}

	return false

}
