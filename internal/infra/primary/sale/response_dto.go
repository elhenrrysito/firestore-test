package sale

type ResponseDTO struct {
	ID            string  `json:"id"`
	OrderNumber   string  `json:"orderNumber"`
	Product       string  `json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Status        string  `json:"status"`
	Date          string  `json:"date"`
	CustomerID    string  `json:"customerID"`
	Total         float64 `json:"total"`
	PaymentMethod string  `json:"paymentMethod"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}
