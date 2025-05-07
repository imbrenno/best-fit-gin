package router

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	// Inicializa as rotas agrupadas
	api := r.Group("/api")
	{
		ProductRoutes(api)
		PingRoutes(api)
	}
}
