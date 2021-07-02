package mysql

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/graph/model"
	"github.com/form3tech-oss/jwt-go"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type LikeRepositoryMysql struct {
	logger *zap.SugaredLogger
}

func NewLikeRepository(l *zap.SugaredLogger) repository.LikeRepository {
	return LikeRepositoryMysql{logger: l}
}

func (likeRepo LikeRepositoryMysql) Create(ctx context.Context, db *ent.Client, restaurantID int) (*model.Like, error) {
	daoLike, err := db.Like.
		Create().
		SetUserID(ctx.Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)).
		SetRestaurantID(restaurantID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	var like *model.Like
	if err := copier.Copy(like, daoLike); err != nil {
		return nil, err
	}

	return like, nil
}
