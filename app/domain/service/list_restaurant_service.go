package service

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

type ListRestaurantService interface {
	Call(ctx context.Context, meshiDB *ent.Client) ([]*ent.Restaurant, error)
}

type listRestaurantService struct {
	logger         *zap.SugaredLogger
	restaurantRepo repository.RestaurantRepository
}

func NewListRestaurantService(l *zap.SugaredLogger, restaurantRepo repository.RestaurantRepository) *listRestaurantService {
	return &listRestaurantService{
		logger:         l,
		restaurantRepo: restaurantRepo,
	}
}

func (s *listRestaurantService) Call(ctx context.Context, meshiDB *ent.Client) ([]*ent.Restaurant, error) {
	return s.restaurantRepo.List(ctx, meshiDB)
}
