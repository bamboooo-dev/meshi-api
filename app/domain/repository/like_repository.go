package repository

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/graph/model"
)

type LikeRepository interface {
	Create(ctx context.Context, db *ent.Client, restaurantID int) (*model.Like, error)
}
