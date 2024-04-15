package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func NoteRoutes(routes *gin.Engine) {
	note := routes.Group("notes")

	note.GET("", controllers.GetNotes)
	note.GET(":note_id", controllers.GetNote)
	note.POST("", controllers.CreateNotes)
	note.PATCH(":note_id", controllers.UpdateNote)
}
