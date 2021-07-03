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
