package property

import "sync"

var (
	firestorePropertyInstance *FirestoreProperty
	onceFirestoreProperty     sync.Once
)

type FirestoreProperty struct {
	Firestore FirestoreConfig `yaml:"firestore"`
}

type FirestoreConfig struct {
	Sales FirestoreCollection `yaml:"sales"`
}

type FirestoreCollection struct {
	Namespace      string `yaml:"namespace"`
	ProjectID      string `yaml:"project-id"`
	CollectionName string `yaml:"collection-name"`
}

func GetFirestoreProperty() *FirestoreProperty {
	onceFirestoreProperty.Do(func() {
		firestorePropertyInstance = &FirestoreProperty{}
	})
	return firestorePropertyInstance
}
