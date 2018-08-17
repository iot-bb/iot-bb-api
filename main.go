package main

import (
	"github.com/iot-bb/iot-bb-api/api"
)

func main() {
	router := api.API()
	router.Run(":8080")
}