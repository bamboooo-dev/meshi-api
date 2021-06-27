package graph

//go:generate go run github.com/99designs/gqlgen
import (
	"github.com/bamboooo-dev/meshi-api/graph/model"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos       []*model.Todo
	restaurants []*model.Restaurant
	logger      *zap.SugaredLogger
}

func NewResolver(l *zap.SugaredLogger) *Resolver {
	return &Resolver{
		logger: l,
	}
}
