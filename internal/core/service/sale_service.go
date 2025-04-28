package service

import (
	"firestore-test/internal/core"
	"firestore-test/internal/core/domain"
	"github.com/google/uuid"
	"time"
)

type SaleService struct {
	persistencePort core.SalePersistencePort
}

func NewSaleService(persistencePort core.SalePersistencePort) *SaleService {
	return &SaleService{persistencePort: persistencePort}
}

func (s *SaleService) Handle(sale *domain.Sale) error {
	// More business logic...
	// ...

	sale.ID = uuid.New().String()
	sale.CreatedAt = time.Now().UTC()
	sale.UpdatedAt = time.Now().UTC()

	return s.persistencePort.Save(*sale)
}

func (s *SaleService) UpdateStatus(orderNumber, status string) error {
	persisted, err := s.persistencePort.FindByOrderNumber(orderNumber)
	if err != nil {
		return err
	}

	if persisted == nil {
		return domain.ErrResourceNotFound
	}

	persisted.Status = status
	persisted.UpdatedAt = time.Now().UTC()

	return s.persistencePort.Update(*persisted)
}
