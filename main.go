package main

import (
	"CitysTempRest/controllers"
	_ "CitysTempRest/docs"
	"CitysTempRest/initializers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// @title           Temperature City Service
// @version         1.0
// @description     A weather service API in Go using Gin framework.

// @contact.name   Davydov Mikhail
// @contact.url    https://github.com/faringet/City_Temp_Rest_Api
// @contact.email  mik.davydov14@gmail.com

// @host      localhost:3000
// @BasePath  /subs
func main() {
	r := gin.Default()

	r.POST("/subs", controllers.SubsCreate)
	r.PUT("/subs/:id", controllers.SubsUpdate)
	r.GET("/subs", controllers.SubsIndex)
	r.GET("/subs/:id", controllers.SubsShow)
	r.DELETE("/subs/:id", controllers.SubsDelete)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
