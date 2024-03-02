package main

import (
	"github.com/Neraverin/gamion-users/initializers"
	"github.com/Neraverin/gamion-users/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectoToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
