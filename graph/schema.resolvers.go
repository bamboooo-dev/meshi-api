package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/graph/generated"
)

func (r *mutationResolver) LikeRestaurant(ctx context.Context, restaurantID int) (*ent.Like, error) {
	service := r.NewLikeRestaurantService()
	return service.Call(ctx, r.db, restaurantID)
}

func (r *queryResolver) Restaurants(ctx context.Context) ([]*ent.Restaurant, error) {
	service := r.NewListRestaurantService()
	return service.Call(ctx, r.db)
}

func (r *queryResolver) NearRestaurants(ctx context.Context) ([]*ent.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FavoriteRestaurants(ctx context.Context) ([]*ent.Restaurant, error) {
	service := r.NewListFavoriteRestaurantService()
	return service.Call(ctx, r.db)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
