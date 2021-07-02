package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	model1 "github.com/bamboooo-dev/meshi-api/app/domain/model"
	"github.com/bamboooo-dev/meshi-api/graph/generated"
	"github.com/bamboooo-dev/meshi-api/graph/model"
)

func (r *mutationResolver) LikeRestaurant(ctx context.Context, restaurantID string) (*model.Like, error) {
	restaurantService := r.NewRestaurantService()

	id, err := strconv.Atoi(restaurantID)
	if err != nil {
		return nil, err
	}
	like, err := restaurantService.Like(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	return like, nil
}

func (r *queryResolver) Restaurants(ctx context.Context) ([]*model1.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) NearRestaurants(ctx context.Context) ([]*model1.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FavoriteRestaurants(ctx context.Context) ([]*model1.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
