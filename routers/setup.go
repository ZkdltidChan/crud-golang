package routers

import (
	"crud-golang/controllers"
	"crud-golang/middlewares"
	"github.com/gin-gonic/gin"
)

// Function to setup routers and router groups
func SetupRouters(app *gin.Engine) {
	v1 := app.Group("/v1")
	{
		v1.GET("/users", controllers.GetAllUsers)
		v1.GET("/users/auth", middlewares.JWTAuthMiddleware(), controllers.GetAllUsers)
		// v1.POST("all", controllers.SendMailUsingMailer)
		v1.POST("/user", controllers.CreateUser)
		v1.POST("/login", controllers.AuthHandler)
		// v1.GET("/ping", controllers.Ping)
		// v1.POST("/user", controllers.CreateUser)
		// v1.GET("/user/:id", controllers.GetUser)
		// v1.GET("/users", controllers.GetUsers)
		// v1.PATCH("/user", controllers.UpdateUser)
		// v1.DELETE("/user/:id", controllers.DeleteUser)
	}
	// Standalone route example
	// app.GET("/ping", controllers.Ping)
}
