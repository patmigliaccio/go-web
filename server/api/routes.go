package api

import (
	"github.com/julienschmidt/httprouter"
)

// Route : a type of object for defining route settings
type Route struct {
	Name   string
	Method string
	Path   string
	Handle httprouter.Handle
}

// Routes : an array of Route settings
type Routes []Route

// AddRoutes : a function to bind the routes to the router
func AddRoutes(root string, r *httprouter.Router) {
	for _, route := range routes {
		r.Handle(route.Method, root+route.Path, route.Handle)
	}
}

var routes = Routes{
	Route{
		"Document",
		"GET",
		"/document/:id",
		DocumentHandler,
	},
}
