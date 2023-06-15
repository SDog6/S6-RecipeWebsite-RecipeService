package main

import (
	"recipeServiceApp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/CraeteRecipe", service.CreateRecipe)
	router.GET("/Recipies", service.GetAllRecipes)
	router.GET("/SingleRecipe", service.GetRecipeByID)

	// Run the router
	router.Run(":9000")

	// kafkaaccess.ConnectAndWriteMessage()
	// kafkaaccess.ConnectAndConsumeMessage()

	//test

}
