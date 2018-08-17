package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Example Api
func Example(router *gin.RouterGroup) {
	router.GET("/example", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Example v1", "status": http.StatusOK})
	})
}
