package api

import (
	"net/http"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

// Document : data model for a document object
// TODO: Move to models directory
type Document struct {
	ID         int    `json:"id"`
	RecordType string `json:"recordType"`
	Method     string `json:"method"`
}

// DocumentHandler : returns a document based on an id it is given
func DocumentHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	document := Document{
		id,
		"Document",
		"GET",
	}

	c.JSON(http.StatusOK, document)
}
