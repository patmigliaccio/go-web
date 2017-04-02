package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

// GetTune : returns a Tune based on an id it is given
func GetTune(c *gin.Context) {
	id := c.Param("id")

	//TODO: Assign the service address elsewhere, not in controller
	addr := viper.GetString("redis.SvcProtocol") + "://" + viper.GetString("redis.SvcHost") + ":" + strconv.Itoa(viper.GetInt("redis.SvcPort"))

	resp, err := http.Get(addr + "/tune/" + id)
	if err != nil {
		log.Println(fmt.Errorf("error on GetTune: %s", err))
		c.JSON(http.StatusBadRequest, err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	c.JSON(http.StatusOK, string(body))
}
