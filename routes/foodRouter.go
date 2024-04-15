package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(routes *gin.Engine) {
	foods := routes.Group("foods")

	foods.GET("", controllers.GetFoods)
	foods.GET(":food_id", controllers.GetFood)
	foods.POST("foods", controllers.CreateFoods)
	foods.PATCH(":food_id", controllers.UpdateFood)
}
