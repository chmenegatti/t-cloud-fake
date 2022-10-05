package routes

import (
	"github.com/chmenegatti/t-cloud-fake/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		tcloud := main.Group("alert")
		{
			tcloud.POST("/", controller.CreateTClouAlert)
		}
	}
	return router
}
