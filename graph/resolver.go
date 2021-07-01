package graph

//go:generate go run github.com/99designs/gqlgen
import (
	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/app/interface/mysql"
	"github.com/bamboooo-dev/meshi-api/ent"
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
	db          *ent.Client
}

func NewResolver(l *zap.SugaredLogger, db *ent.Client) *Resolver {
	return &Resolver{
		logger: l,
		db:     db,
	}
}

func (r *Resolver) NewLikeRepository() repository.LikeRepository {
	return mysql.NewLikeRepository(r.logger)
}
