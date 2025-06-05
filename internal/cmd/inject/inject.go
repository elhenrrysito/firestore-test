//go:build wireinject
// +build wireinject

package inject

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"cloud.google.com/go/firestore"
	"firestore-test/internal/core"
	"firestore-test/internal/core/service"
	configFirestore "firestore-test/internal/infra/config/firestore" // Aliased
	configInstance "firestore-test/internal/infra/config/instance"
	primarySale "firestore-test/internal/infra/primary/sale"
	secondarySale "firestore-test/internal/infra/secondary/persistence/sale"

	// "github.com/gin-gonic/gin" // For GinController and ControllerRunnable - REMOVED TO TEST WIRE
	"github.com/google/wire"
	"gopkg.in/yaml.v3"
)

// --- Start of property structs (mirrors internal/infra/config/property and properties.yml) ---

type ApplicationProperty struct {
	BusinessName string `yaml:"business-name"`
}

type ApplicationConfig struct {
	Application ApplicationProperty `yaml:"application"`
}

type FirestoreCollection struct {
	Namespace      string `yaml:"namespace"`
	ProjectID      string `yaml:"project-id"`
	CollectionName string `yaml:"collection-name"`
}

type FirestoreConfig struct {
	Sales FirestoreCollection `yaml:"sales"`
}

type FirestoreProperty struct {
	Firestore FirestoreConfig `yaml:"firestore"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type ServerProperty struct {
	Server ServerConfig `yaml:"server"`
}

// Combined Properties struct
type Properties struct {
	ApplicationConfig ApplicationConfig `yaml:"application"`
	FirestoreProperty FirestoreProperty `yaml:"firestore"`
	ServerProperty    ServerProperty    `yaml:"server"`
}

// --- End of property structs ---

var (
	onceProperties sync.Once
	props          *Properties
)

// loadProperties reads properties from the YAML file and substitutes environment variables.
func loadProperties() (*Properties, error) {
	onceProperties.Do(func() {
		var p Properties
		absPath, err := filepath.Abs("internal/resources/properties.yml")
		if err != nil {
			log.Fatalf("error getting absolute path for properties: %v", err)
			return // Not strictly necessary due to log.Fatalf, but good practice
		}

		yamlFile, err := ioutil.ReadFile(absPath)
		if err != nil {
			log.Fatalf("error reading properties.yml: %v", err)
			return
		}

		// Substitute environment variables
		expandedYaml := os.ExpandEnv(string(yamlFile))

		err = yaml.Unmarshal([]byte(expandedYaml), &p)
		if err != nil {
			log.Fatalf("error unmarshalling properties.yml: %v", err)
			return
		}
		props = &p
	})
	if props == nil {
		// This case would happen if log.Fatalf didn't actually exit (which it does)
		// or if there was an issue not caught by Fatalf.
		return nil, fmt.Errorf("properties were not loaded")
	}
	return props, nil
}

// --- Property Providers ---

func ProvideProperties() (*Properties, error) {
	return loadProperties()
}

func ProvideServerProperty(p *Properties) *ServerConfig { // Changed to return *ServerConfig
	return &p.ServerProperty.Server
}

func ProvideApplicationProperty(p *Properties) *ApplicationConfig {
	return &p.ApplicationConfig
}

func ProvideFirestoreProperty(p *Properties) *FirestoreProperty {
	return &p.FirestoreProperty
}

// PropertyProviderSet combines all property related providers.
var PropertyProviderSet = wire.NewSet(
	ProvideProperties,
	ProvideServerProperty,
	ProvideApplicationProperty,
	ProvideFirestoreProperty,
)

// --- Firestore Client Provider ---
func ProvideFirestoreClient(fp *FirestoreProperty) (*firestore.Client, error) {
	// Ensure properties are loaded before trying to access them.
	// This explicit call might be removable if ProvideProperties is guaranteed to run first by wire,
	// but it's safer for now or if NewFirestoreClient itself doesn't trigger property loading.
	if _, err := ProvideProperties(); err != nil {
		return nil, fmt.Errorf("failed to load properties before creating firestore client: %w", err)
	}
	// The project ID is nested within the FirestoreProperty struct
	return configFirestore.NewFirestoreClient(fp.Firestore.Sales.ProjectID), nil // Use aliased package
}

// --- Persistence Provider ---
// Depends on the refactored NewRepository
func ProvideSalePersistence(
	client *firestore.Client,
	appConfig *ApplicationConfig,
	firestoreConfig *FirestoreProperty,
) core.SalePersistencePort {
	return secondarySale.NewRepository(
		client,
		firestoreConfig.Firestore.Sales.CollectionName,
		appConfig.Application.BusinessName, // Corrected path to BusinessName
		firestoreConfig.Firestore.Sales.Namespace,
	)
}

// --- Service Provider ---
func ProvideSaleService(persistencePort core.SalePersistencePort) core.SaleUseCaseHandler {
	return service.NewSaleService(persistencePort)
}

// --- Controller Provider ---
func ProvideSaleController(useCase core.SaleUseCaseHandler, persistencePort core.SalePersistencePort) *primarySale.Controller {
	return primarySale.NewController(useCase, persistencePort)
}

// --- GinController (Runnable) Provider ---
// Provider for GinController. Note: instance.GinController and its interface ControllerRunnable will be modified.
func ProvideGinController(saleCtrl *primarySale.Controller, serverProps *ServerConfig) *configInstance.GinController {
	// Adapt this if GinController needs more/different runnables
	runnables := []configInstance.ControllerRunnable{saleCtrl}
	// Pass serverProps.Port (string) instead of the whole ServerConfig object
	return configInstance.NewGinController(runnables, fmt.Sprintf(":%s", serverProps.Port))
}

// AppProviderSet combines all application providers.
var AppProviderSet = wire.NewSet(
	PropertyProviderSet, // Include the previously defined property set
	ProvideFirestoreClient,
	ProvideSalePersistence,
	ProvideSaleService,
	ProvideSaleController,
	ProvideGinController,
)

// --- Injector Function ---
func InitializeApp() (*configInstance.GinController, error) { // Renamed to InitializeApp (Go convention)
	wire.Build(AppProviderSet)
	return nil, nil // Placeholder
}
