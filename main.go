package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Neraverin/gamion-users/initializers"
	"github.com/Neraverin/gamion-users/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectoToDB()
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	router.Run()
}

func getUsers(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Fprintf(os.Stdout, "Can't parse newUser: %v\n", err)
		return
	}

	result := initializers.DB.Create(&newUser)
	if result != nil {
		c.IndentedJSON(http.StatusCreated, newUser)
		return
	}

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Can't create user"})
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.First(&models.User{}, id)
	if result != nil {
		c.IndentedJSON(http.StatusOK, result)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
