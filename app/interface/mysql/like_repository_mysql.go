package mysql

import (
	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"go.uber.org/zap"
)

type LikeRepositoryMysql struct {
	logger *zap.SugaredLogger
}

func NewLikeRepository(l *zap.SugaredLogger) repository.LikeRepository {
	return LikeRepositoryMysql{logger: l}
}
