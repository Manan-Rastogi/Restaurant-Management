package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	users := routes.Group("users")

	users.GET("", controllers.GetUsers)
	users.GET(":user_id", controllers.GetUser)
	users.POST("signup", controllers.Signup)
	users.POST("login", controllers.Login)
}