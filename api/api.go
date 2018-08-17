package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-bb/iot-bb-api/v1"
)
// API pakage
func API() *Engine {
	router := gin.Default()
	apiRouter := router.Group("/api")
	apiv1 := apiRouter.Group("/v1")
	{
		v1.Example(apiv1)
	}
}