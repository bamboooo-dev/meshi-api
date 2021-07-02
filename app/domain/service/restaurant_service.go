package service

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/graph/model"
	"go.uber.org/zap"
)

type RestaurantService interface {
	Like(ctx context.Context, meshiDB *ent.Client, restarurantID int) (*model.Like, error)
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

func (s *restaurantService) Like(ctx context.Context, meshiDB *ent.Client, restarurantID int) (*model.Like, error) {
	like, err := s.likeRepo.Create(ctx, meshiDB, restarurantID)
	if err != nil {
		return nil, err
	}
	return like, nil
}
