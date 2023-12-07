package infrastructure

import (
	"go.uber.org/zap"
)

// LoadLogger ...
func LoadLogger() (*zap.Logger, error) {
	return zap.NewDevelopment()
}
