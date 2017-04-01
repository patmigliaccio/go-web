package api

import (
	"gopkg.in/gin-gonic/gin.v1"
)

// Route : a type of object for defining route settings
type Route struct {
	Name   string
	Method string
	Path   string
	Handle gin.HandlerFunc
}

// Routes : an array of Route settings
type Routes []Route

// AddRoutes : a function to bind the routes to the router
func AddRoutes(root string, r *gin.Engine) {
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
