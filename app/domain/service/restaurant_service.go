package service

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

type RestaurantService interface {
	Like() error
}

type restaurantService struct {
	logger   *zap.SugaredLogger
	likeRepo repository.LikeRepository
}

func NewRestaurantService(l *zap.SugaredLogger, likeRepo repository.LikeRepository) *restaurantService {
	return &restaurantService{
		logger:   l,
		likeRepo: likeRepo,
	}
}

func (s *restaurantService) Like(ctx context.Context, meshiDB *ent.Client) error {
	return nil
}
