package instance

import "firestore-test/internal/infra/primary/sale"

func GetSaleControllerInstance() *sale.Controller {
	return sale.NewController(GetSaleUseCaseInstance(), GetSalePersistenceInstance())
}
