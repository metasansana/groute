package groute

import (
	"net/http"
)

//Route is an interface that is implemented to represent an action that is performed when
//a user enters an application via a specifc uri.
type Route interface {
	Trigger(route string, w http.ResponseWriter, r *http.Request) bool
}
