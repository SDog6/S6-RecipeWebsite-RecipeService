package service

import (
	"recipeServiceApp/dbaccess"
	"recipeServiceApp/model"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRecipe(c *gin.Context) {
	db := dbaccess.ConnectToDb()

	var recipe model.RecipePost
	err := c.BindJSON(&recipe)
	if err != nil {
		log.Fatal("(RegisterUser) c.BindJSON", err)
	}

	query := `INSERT INTO RecipePost (AuthorID,Title,Description,Ingredients,Instructions) VALUES (?,?,?,?,?)`
	res, err := db.Exec(query, recipe.AuthorID, recipe.Title, recipe.Description, recipe.Ingredients, recipe.Instructions)
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
		err := res.Scan(&recipe.ID, &recipe.AuthorID, &recipe.Title, &recipe.Description, &recipe.Ingredients, &recipe.Instructions, &recipe.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		recipies = append(recipies, recipe)
	}

	c.JSON(http.StatusOK, recipies)
}
