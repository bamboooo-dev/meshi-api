package service

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"go.uber.org/zap"
)

type CancelLikeService interface {
	Call(ctx context.Context, meshiDB *ent.Client, restarurantID int) (int, error)
}

type cancelLikeService struct {
	logger   *zap.SugaredLogger
	likeRepo repository.LikeRepository
}

func NewCancelLikeService(l *zap.SugaredLogger, likeRepo repository.LikeRepository) *cancelLikeService {
	return &cancelLikeService{
		logger:   l,
		likeRepo: likeRepo,
	}
}

func (s *cancelLikeService) Call(ctx context.Context, meshiDB *ent.Client, restarurantID int) (int, error) {
	return s.likeRepo.Delete(ctx, meshiDB, restarurantID)
}
