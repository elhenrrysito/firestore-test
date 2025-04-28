package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

func NewFirestoreClient(projectID string) *firestore.Client {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("error trying to create firestore client: %v", err)
	}

	return client
}
