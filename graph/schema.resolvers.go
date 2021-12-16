package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/hirvoin/outfits-server/graph/generated"
	"github.com/hirvoin/outfits-server/graph/model"
	"github.com/hirvoin/outfits-server/internal/garments"
	"github.com/hirvoin/outfits-server/internal/outfits"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateGarment(ctx context.Context, input model.NewGarment) (*model.Garment, error) {
	var garment garments.Garment

	garment.Title = input.Title
	garment.Color = input.Color
	garment.Category = input.Category
	garment.ID = primitive.NewObjectID()

	_, err := garments.CreateGarment(garment)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.Garment{ID: garment.ID.Hex(), Title: garment.Title, Color: garment.Color, Category: garment.Category, WearCount: 0, IsFavorite: false}, nil
}

func (r *mutationResolver) CreateOutfit(ctx context.Context, input model.NewOutfit) (*model.Outfit, error) {
	var outfit outfits.Outfit
	var garmentIds []primitive.ObjectID
	var modelGarments []*model.Garment

	// Get given Garments by id from collection
	dbGarments, getError := garments.GetAll()
	if getError != nil {
		fmt.Println(getError)
		return nil, getError
	}

	// Create slices for garment ids and Garments formatted to model.Garments
	for _, dbGarment := range dbGarments {
		garmentIds = append(garmentIds, dbGarment.ID)
		modelGarments = append(modelGarments, dbGarment.FormatToModel())
	}

	outfit.ID = primitive.NewObjectID()
	outfit.Date = primitive.NewDateTimeFromTime(time.Now())
	outfit.Garments = garmentIds

	// Insert outfit to collection
	_, createError := outfits.CreateOutfit(outfit)

	if createError != nil {
		fmt.Println(createError)
		return nil, createError
	}

	return &model.Outfit{ID: outfit.ID.Hex(), Date: outfit.Date.Time().String(), Garments: modelGarments}, nil
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
	var result []*model.Garment
	dbGarments, _ := garments.GetAll()

	for _, garment := range dbGarments {
		result = append(result, &model.Garment{ID: garment.ID.Hex(), Title: garment.Title, Category: garment.Category, Color: garment.Color, WearCount: garment.WearCount, IsFavorite: garment.IsFavorite})
	}
	return result, nil
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
