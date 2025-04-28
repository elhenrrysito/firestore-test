package sale

import "firestore-test/internal/core/domain"

func toResponseDTO(sale domain.Sale) *ResponseDTO {
	return &ResponseDTO{
		ID:            sale.ID,
		OrderNumber:   sale.OrderNumber,
		Total:         sale.Total,
		Date:          sale.Date,
		Status:        sale.Status,
		PaymentMethod: sale.PaymentMethod,
		CustomerID:    sale.CustomerID,
		Product:       sale.Product,
		Quantity:      sale.Quantity,
		Price:         sale.Price,
		CreatedAt:     sale.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     sale.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func toDomain(sale RequestDTO) *domain.Sale {
	return &domain.Sale{
		OrderNumber:   sale.OrderNumber,
		Total:         sale.Total,
		Status:        sale.Status,
		Date:          sale.Date,
		Price:         sale.Price,
		PaymentMethod: sale.PaymentMethod,
		CustomerID:    sale.CustomerID,
		Product:       sale.Product,
		Quantity:      sale.Quantity,
	}
}
