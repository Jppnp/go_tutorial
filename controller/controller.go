package controller

import (
	"go/tutorial/businessflow"

	"github.com/gin-gonic/gin"
)

func SetupController(router *gin.Engine) {
	router.GET("healthcheck", businessflow.HealthcheckBusineeFlow)
	apiRoutes := router.Group("/api")
	{
		apiRoutes.POST("users/create", businessflow.CreateUserBusinessFlow)
	}
}
