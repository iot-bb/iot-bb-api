package main

import (
	"github.com/iot-bb/iot-bb-api/api"
	"github.com/iot-bb/iot-bb-api/dao"
	// "github.com/iot-bb/iot-bb-api/models"
)
// M map
type M map[string]interface{}

func main() {
	router := api.API()
	db := dao.DAO()
	c := db.C("Test")
	c.Insert(M{"key": "value"})
	print("==>", c.Create)
	router.Run(":8080")
}