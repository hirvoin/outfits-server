package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/hirvoin/outfits-server/graph/generated"
	"github.com/hirvoin/outfits-server/graph/model"
)

func (r *mutationResolver) CreateGarment(ctx context.Context, input model.NewGarment) (*model.Garment, error) {
	var garment model.Garment
	var user model.User
	user.Name = "admin"

	garment.Title = input.Title
	garment.Color = input.Color
	garment.Category = input.Category
	garment.User = &user

	return &garment, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Garments(ctx context.Context) ([]*model.Garment, error) {
	var garments []*model.Garment
	dummyGarment := model.Garment{
		User:       &model.User{Name: "admin"},
		Title:      "Dummy garment",
		Category:   "outerwear",
		Color:      "black",
		WearCount:  0,
		IsFavorite: false,
	}
	garments = append(garments, &dummyGarment)
	return garments, nil
}

func (r *queryResolver) Outfits(ctx context.Context) ([]*model.Outfit, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
