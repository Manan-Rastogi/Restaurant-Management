package main

import (
	"os"

	"github.com/Manan-Rastogi/restaurant-management/database"
	"github.com/Manan-Rastogi/restaurant-management/middleware"
	"github.com/Manan-Rastogi/restaurant-management/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collecton = database.OpenCollection(database.Client, "food")   

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// User Route for Creating a new User.
	routes.UserRoutes(router)

	// Rest of the Routes wi;ll require to authenticate User
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.NoteRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.TableRoutes(router)


	router.Run(":" + port)
}
