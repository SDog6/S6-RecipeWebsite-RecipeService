package main

import (
	"recipeServiceApp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/CraeteRecipe", service.CreateRecipe)
	router.GET("/Recipies", service.GetAllRecipes)
	// router.GET("/products", GetProducts)
	// router.GET("/products/:productId", GetSingleProduct)
	// router.PUT("/products/:productId", UpdateProduct)
	// router.DELETE("/products/:productId", DeleteProduct)

	// Run the router
	router.Run(":8080")

	// kafkaaccess.ConnectAndWriteMessage()
	// kafkaaccess.ConnectAndConsumeMessage()

}
