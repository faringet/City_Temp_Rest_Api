package main

import (
	"CitysTempRest/controllers"
	"CitysTempRest/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/subs", controllers.SubsCreate)
	r.GET("/subs", controllers.SubsIndex)

	r.Run()
}
