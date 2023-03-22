package main

import (
	"crud/crud-module/initializers"
	model "crud/crud-module/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDataBase()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
}
