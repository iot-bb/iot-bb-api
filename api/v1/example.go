package v1

import (
	"net/http"
	"github.com/iot-bb/iot-bb-api/dao"
	"github.com/gin-gonic/gin"
)

// M map
type M map[string]interface{}
var db = dao.DAO()

// Example Api
func Example(router *gin.RouterGroup) {
	router.GET("/example", func (c *gin.Context) {
		println("==> Example v1")
		cc := db.C("Test")
		cc.Insert(M{"api": "v1.0"})
		c.JSON(http.StatusOK, gin.H{"message": "Example v1", "status": http.StatusOK})
	})
}
