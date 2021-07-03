package mysql

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

type RestaurantRepositoryMysql struct {
	logger *zap.SugaredLogger
}

func NewRestaurantRepository(l *zap.SugaredLogger) repository.RestaurantRepository {
	return RestaurantRepositoryMysql{logger: l}
}

func (restaurantRepo RestaurantRepositoryMysql) List(ctx context.Context, db *ent.Client) ([]*ent.Restaurant, error) {
	return db.Restaurant.
		Query().
		All(ctx)
}
