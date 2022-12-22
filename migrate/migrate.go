package main

import (
	"CitysTempRest/initializers"
	"CitysTempRest/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Sub{})
}
