package groute

import (
	"net/http"
)

//Callback coresponds to Callback's signature
type Callback func(w http.ResponseWriter, r *http.Request)

//Route is an inteface used to setup routing.
type Route interface {
	Trigger(route string, w http.ResponseWriter, r *http.Request) bool
}
