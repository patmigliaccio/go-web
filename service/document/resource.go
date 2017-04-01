package document

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"gopkg.in/gin-gonic/gin.v1"
)

// Resource : contains an instance of the Redis Server
type Resource struct {
	Db *redis.Client
}

// GetDocument : returns document by id
func (dr *Resource) GetDocument(c *gin.Context) {
	// TODO: Create document model
	//var document api.Document

	id := c.Param("id")

	document, err := dr.Db.HGetAll("document:" + id).Result()
	if err != nil {
		panic(fmt.Errorf("error retrieving document: %s", err))
	}

	c.JSON(http.StatusOK, document)
}
