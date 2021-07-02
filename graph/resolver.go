package graph

//go:generate go run github.com/99designs/gqlgen
import (
	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/app/domain/service"
	"github.com/bamboooo-dev/meshi-api/app/interface/mysql"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger *zap.SugaredLogger
	db     *ent.Client
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

func (r *Resolver) NewRestaurantService() service.RestaurantService {
	return service.NewRestaurantService(r.logger, r.NewLikeRepository())
}
