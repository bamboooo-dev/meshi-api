package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/graph/generated"
	model1 "github.com/bamboooo-dev/meshi-api/graph/model"
)

func (r *restaurantResolver) Categories(ctx context.Context, obj *ent.Restaurant) ([]*model1.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *restaurantResolver) Location(ctx context.Context, obj *ent.Restaurant) (*model1.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *restaurantResolver) Coordinates(ctx context.Context, obj *ent.Restaurant) (*model1.Coordinates, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *restaurantResolver) Photos(ctx context.Context, obj *ent.Restaurant) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Restaurant returns generated.RestaurantResolver implementation.
func (r *Resolver) Restaurant() generated.RestaurantResolver { return &restaurantResolver{r} }

type restaurantResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *restaurantResolver) ID(ctx context.Context, obj *ent.Restaurant) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
