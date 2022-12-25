package controllers

import (
	"CitysTempRest/initializers"
	"CitysTempRest/models"
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
