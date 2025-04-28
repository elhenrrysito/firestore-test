package instance

import (
	"firestore-test/internal/core"
	"firestore-test/internal/infra/config/firestore"
	"firestore-test/internal/infra/config/property"
	"firestore-test/internal/infra/secondary/persistence/sale"
)

func GetSalePersistenceInstance() core.SalePersistencePort {
	return sale.NewRepository(firestore.NewFirestoreClient(property.GetFirestoreProperty().Firestore.Sales.ProjectID))
}
