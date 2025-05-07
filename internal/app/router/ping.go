package router

import (
	"github.com/gin-gonic/gin"
)

func PingRoutes(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong")
	})
}
