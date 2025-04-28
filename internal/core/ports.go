package core

import "firestore-test/internal/core/domain"

type SalePersistencePort interface {
	Save(sale domain.Sale) error
	FindByOrderNumber(orderNumber string) (*domain.Sale, error)
	Update(sale domain.Sale) error
}
