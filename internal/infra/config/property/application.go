package property

import "sync"

var (
	onceApplicationProperty     sync.Once
	applicationPropertyInstance *ApplicationConfig
)

type ApplicationConfig struct {
	Application ApplicationProperty `yaml:"application"`
}
type ApplicationProperty struct {
	BusinessName string `yaml:"business-name"`
}

func GetApplicationProperty() *ApplicationConfig {
	onceApplicationProperty.Do(func() {
		applicationPropertyInstance = &ApplicationConfig{}
	})
	return applicationPropertyInstance
}
