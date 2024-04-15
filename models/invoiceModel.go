package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	Id             primitive.ObjectID `bson:"_id"`
	InvoiceId      *string             `json:"invoice_id" validate:"required"`
	OrderId        *string             `json:"order_id" validate:"required"`
	PaymentMethod  string             `json:"payment_method" validate:"eq=CARD|eq=CASH|eq=UPI"`
	PaymentStatus  string             `json:"payment_status" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate time.Time          `json:"payment_due_date"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
}
