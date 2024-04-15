package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(routes *gin.Engine) {
	menu := routes.Group("menus")

	menu.GET("", controllers.GetMenus)
	menu.GET(":menu_id", controllers.GetMenu)
	menu.POST("", controllers.CreateMenus)
	menu.PATCH(":menu_id", controllers.UpdateMenu)
}
