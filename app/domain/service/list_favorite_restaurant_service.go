package service

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

type ListFavoriteRestaurantService interface {
	Call(ctx context.Context, meshiDB *ent.Client) ([]*ent.Restaurant, error)
}

type listFavoriteRestaurantService struct {
	logger         *zap.SugaredLogger
	restaurantRepo repository.RestaurantRepository
}

func NewListFavoriteRestaurantService(l *zap.SugaredLogger, restaurantRepo repository.RestaurantRepository) *listFavoriteRestaurantService {
	return &listFavoriteRestaurantService{
		logger:         l,
		restaurantRepo: restaurantRepo,
	}
}

func (s *listFavoriteRestaurantService) Call(ctx context.Context, meshiDB *ent.Client) ([]*ent.Restaurant, error) {
	return s.restaurantRepo.ListByLiker(ctx, meshiDB)
}
