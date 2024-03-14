package businessflow

import (
	"go/tutorial/config"
	"go/tutorial/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HealthcheckBusineeFlow(c *gin.Context) {
	db := config.DB
	healthcheck := new(models.Healthcheck)
	if err := db.Where("id = ?", 1).First(&healthcheck).Error; err != nil {
		panic(err.Error())
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": healthcheck.Message,
		"version": os.Getenv("API_VERSION"),
	})
}
