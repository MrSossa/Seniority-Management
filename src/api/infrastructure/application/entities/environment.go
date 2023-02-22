// Package entities providers entities
package entities

type (
	// Environment type for application environment
	Environment string
	// Environments array of Environment
	Environments []Environment
)

const (
	// ProductionEnvironment constant for production environment
	ProductionEnvironment Environment = "production"
	// SandboxEnvironment constant for sandbox environment
	SandboxEnvironment Environment = "sandbox"
	// DevelopEnvironment constant for develop environment
	DevelopEnvironment Environment = "develop"
	// TestEnvironment constant for test environment
	TestEnvironment Environment = "test"
)

// PossibleEnvironments application possible environments
var PossibleEnvironments = Environments{ProductionEnvironment, SandboxEnvironment, DevelopEnvironment, TestEnvironment}

//

// Contains returns a true bool if a Environment exists in Environments
func (environments Environments) Contains(env Environment) bool {
	for _, environment := range environments {
		if env == environment {
			return true
		}
	}

	return false
}
