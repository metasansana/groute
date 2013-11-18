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

//Get adds a GET route
func (r Router) Get(route Route) Router {

	r.routes[get] = append(r.routes["GET"], route)

	return r

}

//Put adds a PUT route
func (r Router) Put(route Route) Router {

	r.routes[put] = append(r.routes[put], route)

	return r

}

//Post adds a POST route
func (r Router) Post(route Route) Router {

	r.routes[post] = append(r.routes[post], route)

	return r

}

//Delete adds a DELETE route.
func (r Router) Delete(route Route) Router {

	r.routes[del] = append(r.routes[del], route)

	return r

}

//Head adds a HEAD route.
func (r Router) Head(route Route) Router {

	r.routes[head] = append(r.routes[head], route)

	return r

}

//ServeHTTP implements the interface from http.Hanlder. 
func (r Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	for _, aRoute := range r.routes[req.Method] {

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
