package routes

import (
	"github.com/Manan-Rastogi/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(routes *gin.Engine) {
	invoice := routes.Group("invoices")

	invoice.GET("", controllers.GetInvoices)
	invoice.GET(":invoice_id", controllers.GetInvoice)
	invoice.POST("", controllers.CreateInvoices)
	invoice.PATCH(":invoice_id", controllers.UpdateInvoice)
}
