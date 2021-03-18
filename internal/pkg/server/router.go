package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func registerRouter(router *gin.Engine)  {
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
}