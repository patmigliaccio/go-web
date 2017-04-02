package models

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
