package classconnect

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunWebServer() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run()
}
