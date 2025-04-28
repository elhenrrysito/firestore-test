package core

import "firestore-test/internal/core/domain"

type SaleUseCaseHandler interface {
	Handle(sale *domain.Sale) error
	UpdateStatus(orderNumber, status string) error
}
