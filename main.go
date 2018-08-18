package main

import (
	"github.com/iot-bb/iot-bb-api/api"
	// "github.com/iot-bb/iot-bb-api/models"
)

func main() {
	router := api.API()
	router.Run(":8080")
}