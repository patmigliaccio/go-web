package rest

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"gopkg.in/gin-gonic/gin.v1"
)

// Document : data model for a document object
// TODO: Move to models directory
type Document struct {
	ID         string `json:"id"`
	RecordType string `json:"record_type"`
	Value      string `json:"value"`
}

// GetDocument : returns a document based on an id it is given
func GetDocument(c *gin.Context) {
	id := c.Param("id")

	var document Document

	//TODO: Make microservice url dynamic
	resp, err := http.Get("http://localhost:4040/document/" + id)
	if err != nil {
		log.Println(fmt.Errorf("error on GetDocument: %s", err))
		c.JSON(http.StatusBadRequest, err)
	}

	defer resp.Body.Close()

	//TODO: Move unmarshalling into service
	err = json.NewDecoder(resp.Body).Decode(&document)
	if err != nil {
		log.Println(fmt.Errorf("error unmarshalling GetDocument json: %s", err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, document)
}
