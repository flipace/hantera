package lib

import (
	yaml "gopkg.in/yaml.v2"
)

// GetProductConfig : receives a bytestring and returns a ProductConfig
func GetProductConfig(configFile string) ProductConfig {
	config := ProductConfig{}
	err := yaml.Unmarshal(ReadFile(configFile), &config)

	if err != nil {
		panic(err)
	}

	return config
}
