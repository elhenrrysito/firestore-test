package instance

import (
	"firestore-test/internal/core"
	"firestore-test/internal/core/service"
)

func GetSaleUseCaseInstance() core.SaleUseCaseHandler {
	return service.NewSaleService(GetSalePersistenceInstance())
}
