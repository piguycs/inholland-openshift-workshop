package config

import "os"

type ConfigProvider interface {
	GetConfig(name, defaultValue string) string
}

type EnvConfigProvider struct{}

func NewEnvConfigProvider() EnvConfigProvider {
	return EnvConfigProvider{}
}

func (c EnvConfigProvider) GetConfig(name, defaultValue string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultValue
	}

	return value
}
