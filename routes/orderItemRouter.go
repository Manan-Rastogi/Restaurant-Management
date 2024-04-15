package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(routes *gin.Engine) {
	orderItems := routes.Group("orderItemss")

	orderItems.GET("", controllers.GetOrderItems)
	orderItems.GET(":orderItems_id", controllers.GetOrderItems)
	routes.GET("orderItems-order/:order_id", controllers.GetOrderItemsByOrder)
	orderItems.POST("", controllers.CreateOrderItems)
	orderItems.PATCH(":orderItems_id", controllers.UpdateOrderItem)
}
