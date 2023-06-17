package service

import (
	"recipeServiceApp/dbaccess"
	"recipeServiceApp/model"
	"time"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cache "github.com/patrickmn/go-cache"
)

var recipeCache *cache.Cache

func init() {
	// Initialize the cache with a default expiration time of 5 minutes
	recipeCache = cache.New(5*time.Minute, 10*time.Minute)
}

func CreateRecipe(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var recipe model.RecipePost
	err := c.BindJSON(&recipe)
	if err != nil {
		log.Fatal("(RegisterUser) c.BindJSON", err)
	}

	query := `INSERT INTO RecipePost (Author,Title,Description,Picture,Ingredients,Instructions) VALUES (?,?,?,?,?,?)`
	res, err := db.Exec(query, recipe.Author, recipe.Title, recipe.Description, recipe.Picture, recipe.Ingredients, recipe.Instructions)
	if err != nil {
		log.Fatal(err)
	}
	recipe.ID, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, recipe)
}

func GetAllRecipes(c *gin.Context) {

	type RecipeResponse struct {
		Recipies []model.RecipePost `json:"recipies"`
	}

	// Check if the users are already cached
	if cachedRecipies, found := recipeCache.Get("recipies"); found {
		// If the users are cached, return the cached data
		log.Println("From the cache")
		c.JSON(http.StatusOK, RecipeResponse{Recipies: cachedRecipies.([]model.RecipePost)})
		return
	}

	db := dbaccess.ConnectToDb()

	query := "SELECT * FROM RecipePost"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal(err)
	}

	recipies := []model.RecipePost{}
	for res.Next() {
		var recipe model.RecipePost
		err := res.Scan(&recipe.ID, &recipe.Author, &recipe.Title, &recipe.Description, &recipe.Picture,&recipe.Ingredients, &recipe.Instructions, &recipe.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		recipies = append(recipies, recipe)
	}

	// Store the retrieved users in the cache
	recipeCache.Set("recipies", recipies, cache.DefaultExpiration)

	response := RecipeResponse{Recipies: recipies}

	c.JSON(http.StatusOK, response)
}

func GetRecipeByID(c *gin.Context) {

	type RecipeResponse struct {
		Recipies model.RecipePost `json:"recipies"`
	}

	db := dbaccess.ConnectToDb()

// Get the recipe ID from the request body
var requestBody struct {
	ID string `json:"id"`
}

if err := c.ShouldBindJSON(&requestBody); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
}

query := "SELECT * FROM RecipePost WHERE ID = ?"
res, err := db.Query(query, requestBody.ID)
defer res.Close()
if err != nil {
	log.Fatal(err)
}

var recipe model.RecipePost
if res.Next() {
	err := res.Scan(&recipe.ID, &recipe.Author, &recipe.Title, &recipe.Description, &recipe.Picture,&recipe.Ingredients, &recipe.Instructions, &recipe.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}
}

response := RecipeResponse{Recipies: recipe}

c.JSON(http.StatusOK, response)

}

