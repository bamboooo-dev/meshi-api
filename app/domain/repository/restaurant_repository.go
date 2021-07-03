package repository

import (
	"context"

	"github.com/bamboooo-dev/meshi-api/ent"
)

type RestaurantRepository interface {
	List(ctx context.Context, db *ent.Client) ([]*ent.Restaurant, error)
	ListByLiker(ctx context.Context, db *ent.Client) ([]*ent.Restaurant, error)
}
