package main

import (
	"log"
	"github.com/iot-bb/iot-bb-api/api"
	"github.com/iot-bb/iot-bb-api/dao"
	"github.com/iot-bb/iot-bb-api/models"
)

func main() {
	router := api.API()
	db := dao.Dao()
	c = db.C("Test")
	print(c)
	router.Run(":8080")
}