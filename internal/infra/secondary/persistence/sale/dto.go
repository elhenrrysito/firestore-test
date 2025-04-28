package sale

type DTO struct {
	ID            string  `firestore:"id"`
	OrderNumber   string  `firestore:"order_number"`
	Product       string  `firestore:"product"`
	Quantity      int     `firestore:"quantity"`
	Price         float64 `firestore:"price"`
	Date          string  `firestore:"date"`
	Status        string  `firestore:"status"`
	CustomerID    string  `firestore:"customer_id"`
	Total         float64 `firestore:"total"`
	PaymentMethod string  `firestore:"payment_method"`
	CreatedAt     string  `firestore:"created_at"`
	UpdatedAt     string  `firestore:"updated_at"`
}
