package lib

// ProductConfig : config for product
type ProductConfig struct {
	Name         string
	Version      string
	Description  string
	Dependencies map[string]*DependencyConfig
	Steps        Step
}

// DependencyConfig : config for dependency
type DependencyConfig struct {
	Version    string
	Repository string
	Artifacts  string
}

// Step : a step is a command of hantera for which additional commands/functionality can be configured
type Step struct {
	Setup        StepConfiguration
	Update       StepConfiguration
	Dependencies StepConfiguration
}

// StepConfiguration : contains information about pre/post/override commands for a step
type StepConfiguration struct {
	Pre      []string
	Override []string
	Post     []string
}
