package property

import "sync"

var (
	serverPropertyInstance *ServerProperty
	serverPropertyOnce     sync.Once
)

type ServerProperty struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func GetServerProperty() *ServerProperty {
	serverPropertyOnce.Do(func() {
		serverPropertyInstance = &ServerProperty{}
	})
	return serverPropertyInstance
}
