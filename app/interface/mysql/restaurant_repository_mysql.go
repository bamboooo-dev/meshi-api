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

type RestaurantRepositoryMysql struct {
	logger *zap.SugaredLogger
}

func NewRestaurantRepository(l *zap.SugaredLogger) repository.RestaurantRepository {
	return RestaurantRepositoryMysql{logger: l}
}

func (restaurantRepo RestaurantRepositoryMysql) List(ctx context.Context, client *ent.Client) ([]*ent.Restaurant, error) {
	return client.Restaurant.
		Query().
		All(ctx)
}

func (restaurantRepo RestaurantRepositoryMysql) ListByLiker(ctx context.Context, client *ent.Client) ([]*ent.Restaurant, error) {
	return client.Restaurant.
		Query().
		Where(
			restaurant.HasLikesWith(
				like.UserID(ctx.Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"].(string)),
			),
		).
		All(ctx)
}
