package router

import (
	"github.com/gin-gonic/gin"
	c "github.com/imbrenno/best-fit-gin/internal/app/controller"
)

func ProductRoutes(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.POST("", c.ProductCreate)      // Create
		products.GET("", c.ProductList)         // Read all
		products.GET(":id", c.ProductId)        // Read One
		products.PUT(":id", c.ProductUpdate)    // Update
		products.DELETE(":id", c.ProductDelete) // Delete
	}
}
