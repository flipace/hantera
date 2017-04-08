package lib

// DependencyConfig : config for dependency
type DependencyConfig struct {
	Version    string
	Repository string
	Artifacts  string
}

// ProductConfig : config for product
type ProductConfig struct {
	Name         string
	Version      string
	Description  string
	Dependencies map[string]*DependencyConfig
}
