package grpcApi

import (
	"context"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/pbx"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (server *Server) CreateRecipe(ctx context.Context, req *pb.CreateRecipeRequest) (*pb.CreateRecipeResponse, error) {
	authorizeUserPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateCreateRecipeRequest(req); err != nil {
		return nil, err
	}

	recipe := &models.Recipe{
		Title:       req.Recipe.Title,
		Description: null.NewString(req.Recipe.GetDescription(), true),
		Directions:  null.NewString(req.Recipe.GetInstructions(), true),
		AuthorID:    authorizeUserPayload.ID.String(),
	}

	for _, ingredient := range req.Recipe.Ingredients {

		if ingredient.Id == 0 {
			ingredientDb := &models.Ingredient{
				Name: ingredient.Name,
			}

			err := ingredientDb.Insert(ctx, server.config.DB, boil.Infer())
			if err != nil {
				return nil, err
			}

			ingredient.Id = int64(ingredientDb.ID)
		}

		recipeIngredient := &models.RecipeIngredient{
			RecipeID:     recipe.ID,
			Quantity:     float64(ingredient.Quantity),
			IngredientID: int(ingredient.Id),
		}
		recipe.AddRecipeIngredients(ctx, server.config.DB, true, recipeIngredient)
	}

	err = recipe.Insert(ctx, server.config.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.CreateRecipeResponse{
		Recipe: pbx.DbRecipeToProto(recipe),
	}, nil
}

func validateCreateRecipeRequest(req *pb.CreateRecipeRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Recipe,
			validation.Required,
			validation.By(
				func(value interface{}) error {
					return validateRecipe(value.(*pb.Recipe))
				},
			),
		),
	)
}

func validateRecipe(recipe *pb.Recipe) error {
	return validation.ValidateStruct(recipe,
		validation.Field(&recipe.Title,
			validation.Required,
			validation.Length(1, 100),
		),
		validation.Field(&recipe.Description,
			validation.Required,
			validation.Length(1, 250),
		),
		validation.Field(&recipe.Instructions,
			validation.Required,
		),
		validation.Field(&recipe.Ingredients,
			validation.Required,
			validation.Length(1, 0), // Specify the maximum allowed length for ingredients if needed
		),
	)
}
