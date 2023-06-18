package main

import (
	"recipeServiceApp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	recipeTable := "RecipePost"
	router.POST("/CraeteRecipe", func(c *gin.Context) {
		service.CreateRecipe(c, recipeTable)
	})
	router.GET("/Recipies", func(c *gin.Context) {
		service.GetAllRecipes(c, recipeTable)
	})
	router.POST("/SingleRecipe", func(c *gin.Context) {
		service.GetRecipeByID(c, recipeTable)
	})

	// Run the router
	router.Run(":9000")

	// kafkaaccess.ConnectAndWriteMessage()
	// kafkaaccess.ConnectAndConsumeMessage()

	//test

}
