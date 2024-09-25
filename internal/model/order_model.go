package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	Id          uint64      `json:"id"`
	CustomerId  uuid.UUID   `json:"customerId"`
	OrderItems  []OrderItem `json:"orderItems"`
	CreatedAt   *time.Time  `json:"createdAt"`
	ShippedAt   *time.Time  `json:"shippedAt"`
	CompletedAt *time.Time  `json:"completedAt"`
}

type OrderItem struct {
	Id       uuid.UUID `json:"id"`
	Quantity uint      `json:"quantity"`
	Price    uint      `json:"price"`
}
