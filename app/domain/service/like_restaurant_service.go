package service

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

type LikeRestaurantService interface {
	Call(ctx context.Context, meshiDB *ent.Client, restarurantID string) (*ent.Like, error)
}

type likeRestaurantService struct {
	logger   *zap.SugaredLogger
	likeRepo repository.LikeRepository
}

func NewLikeRestaurantService(l *zap.SugaredLogger, likeRepo repository.LikeRepository) *likeRestaurantService {
	return &likeRestaurantService{
		logger:   l,
		likeRepo: likeRepo,
	}
}

func (s *likeRestaurantService) Call(ctx context.Context, meshiDB *ent.Client, restarurantID string) (*ent.Like, error) {
	return s.likeRepo.Create(ctx, meshiDB, restarurantID)
}
