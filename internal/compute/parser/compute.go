package parser

import "go.uber.org/zap"

type Compute struct {
	logger *zap.Logger
}

func NewCompute(logger *zap.Logger) *Compute {
	return &Compute{
		logger: logger,
	}
}
