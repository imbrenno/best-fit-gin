package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imbrenno/best-fit-gin/internal/app/database"
	"github.com/imbrenno/best-fit-gin/internal/app/model"
	"github.com/imbrenno/best-fit-gin/internal/app/utils/enum"
)

// ProductCreate handles the creation of a new product
func ProductCreate(c *gin.Context) {
	var product model.Product

	// Faz o bind do JSON recebido para o objeto product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Product Details: %+v\n", c.Request.Body)
	if err := enum.ValidadeProductType(c.Param("product_type")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Salva o produto no banco de dados
	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully",
		"id": product.ID})
}

func ProductList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of products"})
}

func ProductId(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Product details", "id": id})
}

func ProductUpdate(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "id": id})
}

func ProductDelete(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully", "id": id})
}
