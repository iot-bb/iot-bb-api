package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iot-bb/iot-bb-api"
)

func main() {
	router := api.API()
	router.Run(":8080")
}