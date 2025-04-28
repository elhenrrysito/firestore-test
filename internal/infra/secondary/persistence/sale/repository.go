package sale

import (
	"cloud.google.com/go/firestore"
	"context"
	"firestore-test/internal/core/domain"
	"firestore-test/internal/infra/config/property"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Repository struct {
	client         *firestore.Client
	collectionName string
}

func NewRepository(client *firestore.Client) *Repository {
	return &Repository{client: client, collectionName: property.GetFirestoreProperty().Firestore.Sales.CollectionName}
}

func (r *Repository) Save(sale domain.Sale) error {
	saleDTO := toDTO(sale)

	_, err := r.client.Collection(property.GetApplicationProperty().Application.BusinessName).
		Doc(property.GetFirestoreProperty().Firestore.Sales.Namespace).
		Collection(r.collectionName).
		Doc(sale.OrderNumber).
		Create(context.Background(), saleDTO)
	if err != nil {
		return fmt.Errorf("could not save sale with order number %v: %w", sale.OrderNumber, err)
	}

	return nil
}

func (r *Repository) FindByOrderNumber(orderNumber string) (*domain.Sale, error) {
	doc, err := r.client.Collection(property.GetApplicationProperty().Application.BusinessName).
		Doc(property.GetFirestoreProperty().Firestore.Sales.Namespace).
		Collection(r.collectionName).
		Doc(orderNumber).
		Get(context.Background())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("error trying to find sale with order number %v: %w", orderNumber, err)
	}

	if !doc.Exists() {
		return nil, nil
	}

	var saleDTO DTO

	err = doc.DataTo(&saleDTO)
	if err != nil {
		return nil, fmt.Errorf("error trying to convert sale data to DTO for order number %s: %w", orderNumber, err)
	}

	sale, err := toDomain(saleDTO)
	if err != nil {
		return nil, fmt.Errorf("error trying to convert sale DTO to domain for order number %s: %w", orderNumber, err)
	}

	return sale, nil
}

func (r *Repository) Update(sale domain.Sale) error {
	saleDTO := toDTO(sale)

	_, err := r.client.Collection(property.GetApplicationProperty().Application.BusinessName).
		Doc(property.GetFirestoreProperty().Firestore.Sales.Namespace).
		Collection(r.collectionName).
		Doc(sale.OrderNumber).
		Set(context.Background(), saleDTO)

	if err != nil {
		return fmt.Errorf("could not update sale with order number %v: %w", sale.OrderNumber, err)
	}

	return nil
}
