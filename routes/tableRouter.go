package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(routes *gin.Engine) {
	table := routes.Group("tables")

	table.GET("", controllers.GetTables)
	table.GET(":table_id", controllers.GetTable)
	table.POST("", controllers.CreateTables)
	table.PATCH(":table_id", controllers.UpdateTable)
}
