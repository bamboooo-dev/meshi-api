package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/bamboooo-dev/meshi-api/app/domain/service"
	"github.com/bamboooo-dev/meshi-api/graph/generated"
	"github.com/bamboooo-dev/meshi-api/graph/model"
	jwt "github.com/form3tech-oss/jwt-go"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) LikeRestaurant(ctx context.Context, restaurantID string) (*model.Like, error) {

	restaurantService := service.NewRestaurantService(r.logger, r.NewLikeRepository())
	if err := restaurantService.Like(ctx, r.db); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	fmt.Println(ctx.Value("user").(*jwt.Token).Claims.(jwt.MapClaims)["sub"])
	return r.todos, nil
}

func (r *queryResolver) Restaurants(ctx context.Context) ([]*model.Restaurant, error) {
	return r.restaurants, nil
}

func (r *queryResolver) NearRestaurants(ctx context.Context) ([]*model.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FavoriteRestaurants(ctx context.Context) ([]*model.Restaurant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
