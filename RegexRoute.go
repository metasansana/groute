package groute

import (
	"net/http"
	"regexp"
)

//RegexRoute provides regex routing support.
type RegexRoute struct {
	path   *regexp.Regexp
	action Callback
}

//Trigger implements the method from the Route interface.
func (self *RegexRoute) Trigger(route string, w http.ResponseWriter, r *http.Request) bool {

	if params := self.path.FindStringSubmatch(route); len(params) > 0 {

		self.action(w, r)

		return true

	}
	return false
}

//NewRegexRoute constructor.
func NewRegexRoute(path string, callback Callback) Route {

	return &RegexRoute{regexp.MustCompile(path), callback}

}
