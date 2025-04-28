package domain

import "time"

type Sale struct {
	ID            string
	OrderNumber   string
	Product       string
	Quantity      int
	Price         float64
	Date          string
	Status        string
	CustomerID    string
	Total         float64
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
