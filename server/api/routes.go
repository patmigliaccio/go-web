package api

import (
	"github.com/julienschmidt/httprouter"
)

// ROOT : sets the default root path for api calls
// TODO: Move to environment config file
const ROOT string = "/api"

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
func AddRoutes(r *httprouter.Router) {
	for _, route := range routes {
		r.Handle(route.Method, ROOT+route.Path, route.Handle)
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
