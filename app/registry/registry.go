package registry

import "go.uber.org/zap"

type Registry interface{}

type registry struct {
	l *zap.SugaredLogger
}

func NewRegistry(l *zap.SugaredLogger) Registry {
	return &registry{l}
}
