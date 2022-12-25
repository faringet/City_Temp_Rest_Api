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
	r.PUT("/subs/:id", controllers.SubsUpdate)
	r.GET("/subs", controllers.SubsIndex)
	r.GET("/subs/:id", controllers.SubsShow)

	r.Run()
}
