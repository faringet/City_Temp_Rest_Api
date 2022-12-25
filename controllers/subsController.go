package controllers

import (
	"CitysTempRest/initializers"
	"CitysTempRest/models"
	"CitysTempRest/weather"
	"github.com/gin-gonic/gin"
)

func SubsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		City        string
		Temperature float64
	}

	c.Bind(&body)

	// Create a sub
	sub := models.Sub{City: body.City}

	result := initializers.DB.Create(&sub)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"sub": sub,
	})

	//Put temperature to users City
	var a = weather.GetTemperature(weather.GetCityID(weather.ReturnMap(), sub.City))
	initializers.DB.First(&sub, sub.ID)
	initializers.DB.Model(&sub).Updates(models.Sub{Temperature: a})

}

func SubsIndex(c *gin.Context) {
	// Get the subs
	var subs []models.Sub
	initializers.DB.Find(&subs)

	// Respond with them
	c.JSON(200, gin.H{
		"sub": subs,
	})
}

func SubsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the subs
	var sub models.Sub
	initializers.DB.First(&sub, id)

	// Respond with them
	c.JSON(200, gin.H{
		"sub": sub,
	})
}

func SubsUpdate(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		City        string
		Temperature float64
	}

	c.Bind(&body)

	// Find the sub were updating
	var sub models.Sub
	initializers.DB.First(&sub, id)

	// Update it
	var a = weather.GetTemperature(weather.GetCityID(weather.ReturnMap(), sub.City))
	initializers.DB.Model(&sub).Updates(models.Sub{City: body.City, Temperature: a})

	// Respond with it
	c.JSON(200, gin.H{
		"sub": sub,
	})
}

func SubsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Delete the subs
	initializers.DB.Delete(&models.Sub{}, id)

	// Respond
	c.Status(200)

}
