package main

import (
	"go/tutorial/config"
	"go/tutorial/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	app := gin.Default()

	app.Use(cors.Default())

	controller.SetupController(app)
	app.Run(":8080")
}
