package api

import (
  "github.com/gin-gonic/gin"
  "github.com/iot-bb/iot-bb-api/api/v1"
)

// API package
func API() *gin.Engine {
  router := gin.Default()
  apiRouter := router.Group("/api")
  apiv1 := apiRouter.Group("/v1")
    {
      v1.Users(apiv1)
    }
  return router
}