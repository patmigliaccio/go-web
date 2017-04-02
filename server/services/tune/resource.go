package tune

import (
	"fmt"
	"net/http"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/patmigliaccio/go-web/server/models"
	"gopkg.in/gin-gonic/gin.v1"
)

// Resource : contains an instance of the Redis Server
type Resource struct {
	Db *redis.Client
}

// GetTune : returns tune by id
func (tr *Resource) GetTune(c *gin.Context) {
	id := c.Param("id")

	tune, err := tr.Db.HGetAll("tune:" + id).Result()
	if err != nil {
		panic(fmt.Errorf("error retrieving tune: %s", err))
	}

	c.JSON(http.StatusOK, tune)
}

// AddTune : creates a new tune and adds it to database
func (tr *Resource) AddTune(c *gin.Context) {
	var tune models.Tune

	err := c.Bind(&tune)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, tune)
		return
	}

	tune.ID = uuid.New().String()

	tr.Db.HMSet("tune:"+tune.ID, structs.Map(tune))
	c.JSON(http.StatusOK, tune)
}
