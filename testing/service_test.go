package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"recipeServiceApp/model"
	"recipeServiceApp/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestCreateRecipe(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create the request body
	var recipe = model.RecipePost{
		Author:       "John Doe",
		Title:        "Delicious Chocolate Cake",
		Description:  "A mouthwatering chocolate cake recipe",
		Picture:      "https://example.com/cake.jpg",
		Ingredients:  "Flour, sugar, cocoa powder, eggs, milk",
		Instructions: "1. Preheat the oven. 2. Mix dry ingredients. 3. Beat eggs and sugar. 4. Combine dry and wet ingredients. 5. Bake for 30 minutes.",
	}

	payload, err := json.Marshal(recipe)
	if err != nil {
		t.Fatal("Failed to marshal JSON payload:", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/create-recipe", bytes.NewBuffer(payload))

	// Create a test response recorder
	recorder := httptest.NewRecorder()
	// Call the CreateRecipe handler function
	table := "RecipePost_Testing" // Provide the table name
	router.POST("/create-recipe", func(c *gin.Context) {
		service.CreateRecipe(c, table)
	})
	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAllRecipes(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create a test request
	request := httptest.NewRequest(http.MethodGet, "/recipes", nil)

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the GetAllRecipes handler function
	table := "RecipePost_Testing" // Provide the table name
	router.GET("/recipes", func(c *gin.Context) {
		service.GetAllRecipes(c, table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
	// Assert the response body or other details as needed
}

func TestGetRecipeByID(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create the request body
	requestBody := map[string]int64{
		"id": 1,
	}
	payload, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal("Failed to marshal JSON payload:", err)
	}
	request := httptest.NewRequest(http.MethodGet, "/recipe", bytes.NewBuffer(payload))

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the GetRecipeByID handler function
	table := "RecipePost_Testing" // Provide the table name
	router.GET("/recipe", func(c *gin.Context) {
		service.GetRecipeByID(c, table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
	// Assert the response body or other details as needed
}
