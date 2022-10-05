package controller

import (
	"github.com/chmenegatti/t-cloud-fake/database"
	"github.com/chmenegatti/t-cloud-fake/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTClouAlert(c *gin.Context) {
	db := database.GetDatabase()

	var alert models.Alert

	err := c.ShouldBindJSON(&alert)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON" + err.Error(),
		})

		return
	}

	alert.ID = uuid.New().String()

	err = db.Create(&alert).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create alert" + err.Error(),
		})

		return
	}

	c.JSON(200, alert)

}
