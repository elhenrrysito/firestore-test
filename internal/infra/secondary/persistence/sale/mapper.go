package sale

import (
	"firestore-test/internal/core/domain"
	"time"
)

func toDTO(sale domain.Sale) *DTO {
	return &DTO{
		ID:            sale.ID,
		OrderNumber:   sale.OrderNumber,
		Product:       sale.Product,
		Quantity:      sale.Quantity,
		Price:         sale.Price,
		Date:          sale.Date,
		Status:        sale.Status,
		CustomerID:    sale.CustomerID,
		Total:         sale.Total,
		PaymentMethod: sale.PaymentMethod,
		CreatedAt:     sale.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:     sale.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func toDomain(dto DTO) (*domain.Sale, error) {
	createdAt, err := time.Parse("2006-01-02 15:04:05", dto.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := time.Parse("2006-01-02 15:04:05", dto.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &domain.Sale{
		ID:            dto.ID,
		OrderNumber:   dto.OrderNumber,
		Product:       dto.Product,
		Quantity:      dto.Quantity,
		Price:         dto.Price,
		Status:        dto.Status,
		Date:          dto.Date,
		CustomerID:    dto.CustomerID,
		Total:         dto.Total,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		PaymentMethod: dto.PaymentMethod,
	}, nil
}
