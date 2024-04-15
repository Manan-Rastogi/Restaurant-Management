package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(routes *gin.Engine) {
	order := routes.Group("orders")

	order.GET("", controllers.GetOrders)
	order.GET(":order_id", controllers.GetOrder)
	order.POST("", controllers.CreateOrders)
	order.PATCH(":order_id", controllers.UpdateOrder)
}
