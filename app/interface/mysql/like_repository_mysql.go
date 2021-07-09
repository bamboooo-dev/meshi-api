package mysql

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/app/domain/repository"
	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/ent/like"
	"github.com/bamboooo-dev/meshi-api/ent/restaurant"
	"github.com/form3tech-oss/jwt-go"
	"go.uber.org/zap"
)

type LikeRepositoryMysql struct {
	logger *zap.SugaredLogger
}

func NewLikeRepository(l *zap.SugaredLogger) repository.LikeRepository {
	return LikeRepositoryMysql{logger: l}
}

func (likeRepo LikeRepositoryMysql) Create(ctx context.Context, db *ent.Client, restaurantID string) (*ent.Like, error) {
	return db.Like.
		Create().
		SetUserID(ctx.Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)).
		SetRestaurantID(restaurantID).
		Save(ctx)
}

func (likeRepo LikeRepositoryMysql) Delete(ctx context.Context, db *ent.Client, restaurantID string) (int, error) {
	return db.Like.
		Delete().
		Where(
			like.UserID(ctx.Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)),
			like.HasRestaurantWith(restaurant.ID(restaurantID)),
		).
		Exec(ctx)
}
