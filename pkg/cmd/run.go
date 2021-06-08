package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/kayes-shawon/go-gin/pkg/api"
	"github.com/kayes-shawon/go-gin/pkg/middleware"
)

func Restricted(app *gin.Engine)  {
	authGroup := app.Group("api/v1", middleware.Auth)
	studentGroup := authGroup.Group("student")
	{
		studentGroup.POST("/create", api.StudentCreate)
	}
}

func UnRestricted(app *gin.Engine)  {
	app.Group("/user", func(c *gin.Context) {
		
	})
}

func RunServer() {
	app := gin.New()

	UnRestricted(app)

	app.Run(":8080")
}