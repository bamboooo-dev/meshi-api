package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bamboooo-dev/meshi-api/ent"
	"github.com/bamboooo-dev/meshi-api/graph/generated"
)

func (r *likeResolver) RestaurantID(ctx context.Context, obj *ent.Like) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Like returns generated.LikeResolver implementation.
func (r *Resolver) Like() generated.LikeResolver { return &likeResolver{r} }

type likeResolver struct{ *Resolver }
