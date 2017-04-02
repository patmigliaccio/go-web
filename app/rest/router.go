package rest

import (
	"github.com/patmigliaccio/go-web/app/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// Routes : an array of Route settings
type Routes []models.Route

// AddRoutes : a function to bind the routes to the router
func AddRoutes(root string, r *gin.Engine) {
	for _, route := range routes {
		r.Handle(route.Method, root+route.Path, route.Handle)
	}
}

var routes = Routes{
	models.Route{
		"Document",
		"GET",
		"/document/:id",
		GetDocument,
	},
}
