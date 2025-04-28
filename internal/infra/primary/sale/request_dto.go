package sale

type RequestDTO struct {
	OrderNumber   string  `json:"orderNumber"`
	Product       string  `json:"product"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	Status        string  `json:"status"`
	Date          string  `json:"date"`
	CustomerID    string  `json:"customerID"`
	Total         float64 `json:"total"`
	PaymentMethod string  `json:"paymentMethod"`
}

type UpdateRequestDTO struct {
	OrderNumber string `json:"orderNumber"`
	Status      string `json:"status"`
}
