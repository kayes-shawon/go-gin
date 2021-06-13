package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/kayes-shawon/go-gin/pkg/api"
)

func Restricted(app *gin.Engine)  {
	//authGroup := app.Group("api/v1", middleware.Auth)
	//studentGroup := authGroup.Group("student")
	{
		//studentGroup.POST("/create", api.CreateStudent)
		//studentGroup.GET("/list", api.StudentList)
		//studentGroup.GET("/:id", api.StudentDetails)
		//studentGroup.PUT("/update/:id", api.UpdateStudent)
		//studentGroup.PUT("/delete/:id", api.DeleteStudent)
	}
}

func UnRestricted(app *gin.Engine)  {
	userGroup := app.Group("/user")
	{
		userGroup.POST("/create", api.CreateUser)
		userGroup.POST("/login", api.UserLogin)
		userGroup.POST("/login-refresh", api.UserLoginRefresh)
	}

}

func RunServer() {
	app := gin.New()

	UnRestricted(app)
	//Restricted(app)

	app.Run(":8080")
}