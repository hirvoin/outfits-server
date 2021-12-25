package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
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
	garment.ImageUri = input.ImageURI

	_, err := garments.CreateGarment(garment)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.Garment{ID: garment.ID.Hex(), Title: garment.Title, Color: garment.Color, Category: garment.Category, ImageURI: garment.ImageUri, WearCount: 0, IsFavorite: false}, nil
}

func (r *mutationResolver) CreateOutfit(ctx context.Context, input model.NewOutfit) (*model.Outfit, error) {
	var outfit outfits.Outfit
	var garmentObjectIds []primitive.ObjectID
	var modelGarments []*model.Garment

	// Convert garment string ids to ObjectIds
	for _, stringId := range input.Garments {
		objId, _ := primitive.ObjectIDFromHex(stringId)
		garmentObjectIds = append(garmentObjectIds, objId)
	}

	// Get given Garments by ObjectIds from collection
	dbGarments, getError := garments.GetGarmentsByIds(garmentObjectIds)
	if getError != nil {
		fmt.Println(getError)
		return nil, getError
	}

	// Convert dbGarments to  GraphQL schema model.Garments
	for _, dbGarment := range dbGarments {
		modelGarments = append(modelGarments, dbGarment.FormatToModel())
	}

	outfit.ID = primitive.NewObjectID()
	outfit.Date = primitive.NewDateTimeFromTime(time.Now())
	outfit.Garments = garmentObjectIds

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

func (r *mutationResolver) UpdateGarment(ctx context.Context, input *model.UpdatedGarment) (*model.Garment, error) {
	var garment garments.Garment

	replacedGarment, err := garments.GetGarmentById(input.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	garment.ID = replacedGarment.ID
	garment.IsFavorite = replacedGarment.IsFavorite
	garment.WearCount = replacedGarment.WearCount
	garment.Title = input.Title
	garment.Color = input.Color
	garment.Category = input.Category
	garment.ImageUri = input.ImageURI

	updatedGarment, err := garments.EditGarment(garment)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.Garment{ID: updatedGarment.ID.Hex(), Title: updatedGarment.Title, Color: updatedGarment.Color, Category: updatedGarment.Category, ImageURI: updatedGarment.ImageUri, WearCount: updatedGarment.WearCount, IsFavorite: updatedGarment.IsFavorite}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Garments(ctx context.Context, category *string, id *string) ([]*model.Garment, error) {
	var result []*model.Garment

	if id != nil {
		garment, err := garments.GetGarmentById(*id)
		if err != nil {
			return result, errors.New("No garments found with given id.")
		}
		result = append(result, garment.FormatToModel())
		return result, nil
	}

	if category != nil && *category != "outerwear" && *category != "tops" && *category != "bottoms" && *category != "footwear" {
		return result, errors.New("Invalid category: " + *category)
	}

	dbGarments, _ := garments.GetAll()

	for _, garment := range dbGarments {
		if category == nil || *category == garment.Category {
			result = append(result, &model.Garment{ID: garment.ID.Hex(), Title: garment.Title, Category: garment.Category, Color: garment.Color, WearCount: garment.WearCount, IsFavorite: garment.IsFavorite, ImageURI: garment.ImageUri})
		}
	}
	return result, nil
}

func (r *queryResolver) Outfits(ctx context.Context) ([]*model.Outfit, error) {
	var result []*model.Outfit

	dbOutfits, outfitsError := outfits.GetAll()
	if outfitsError != nil {
		fmt.Println(outfitsError)
		return nil, outfitsError
	}

	for _, outfit := range dbOutfits {
		var modelGarments []*model.Garment

		dbGarments, garmentError := garments.GetGarmentsByIds(outfit.Garments)
		if garmentError != nil {
			fmt.Println(garmentError)
			return nil, garmentError
		}

		// Create slices for outfit.Garments formatted to model.Garments
		for _, dbGarment := range dbGarments {
			modelGarments = append(modelGarments, dbGarment.FormatToModel())
		}

		result = append(result, &model.Outfit{ID: outfit.ID.Hex(), Garments: modelGarments, Date: outfit.Date.Time().String()})
	}
	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
