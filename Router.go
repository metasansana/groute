package groute

import (
	"net/http"
)

const (
	get  = "GET"
	put  = "PUT"
	post = "POST"
	del  = "DELETE"
	head = "HEAD"
)

type list map[string][]Route

//Router is the object that sets up routing.
type Router struct {
	routes list
}

//Regex is a factory method for regex route support.
func (r *Router) Regex(regex string, callback Callback) Route {

	return NewRegexRoute(regex, callback)
}

//Static is a factory method for static route support.
func (self *Router) Static(path string, callback Callback) Route {

	return NewStaticRoute(path, callback)
}

//Get adds a GET route
func (self *Router) Get(route Route) *Router {

	self.routes[get] = append(self.routes["GET"], route)

	return self

}

//Put adds a PUT route
func (self *Router) Put(route Route) *Router {

	self.routes[put] = append(self.routes[put], route)

	return self

}

//Post adds a POST route
func (self *Router) Post(route Route) *Router {

	self.routes[post] = append(self.routes[post], route)

	return self

}

//Delete adds a DELETE route.
func (self *Router) Delete(route Route) *Router {

	self.routes[del] = append(self.routes[del], route)

	return self

}

//Head adds a HEAD route.
func (self *Router) Head(route Route) *Router {

	self.routes[head] = append(self.routes[head], route)

	return self

}

//ServeHTTP implements the interface from http.Hanlder.
func (self *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	for _, aRoute := range self.routes[req.Method] {

		if result := aRoute.Trigger(req.URL.Path, res, req); result {

			return

		}

	}

}

//NewRouter constructs a new Router object.
func NewRouter() *Router {

	routes := make(list)
	routes["GET"] = make([]Route, 0)
	routes["PUT"] = make([]Route, 0)
	routes["POST"] = make([]Route, 0)
	routes["DELETE"] = make([]Route, 0)
	routes["HEAD"] = make([]Route, 0)

	return &Router{routes}

}
