package repository

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/ent"
)

type LikeRepository interface {
	Create(ctx context.Context, db *ent.Client, restaurantID int) (*ent.Like, error)
	Delete(ctx context.Context, db *ent.Client, restaurantID int) (int, error)
}
