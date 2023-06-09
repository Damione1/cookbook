package grpcApi

import (
	"context"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/pbx"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (server *Server) GetRecipe(ctx context.Context, req *pb.GetRecipeRequest) (*pb.GetRecipeResponse, error) {
	if err := validateGetRecipeRequest(req); err != nil {
		return nil, err
	}
	recipe, err := models.Recipes(models.RecipeWhere.ID.EQ(int(req.Id))).One(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	return &pb.GetRecipeResponse{
		Recipe: pbx.DbRecipeToProto(recipe),
	}, nil
}

func (server *Server) ListRecipes(ctx context.Context, req *pb.ListRecipesRequest) (*pb.ListRecipesResponse, error) {
	if err := validateListRecipesRequest(req); err != nil {
		return nil, err
	}
	recipes, err := models.Recipes().All(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	response := &pb.ListRecipesResponse{}
	response.Recipes = make([]*pb.Recipe, len(recipes))

	for i, recipe := range recipes {
		response.Recipes[i] = pbx.DbRecipeToProto(recipe)
	}

	return response, nil
}

func (server *Server) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*pb.DeleteRecipeResponse, error) {
	authorizeUserPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateDeleteRecipeRequest(req); err != nil {
		return nil, err
	}

	id := req.Id
	recipe, err := models.Recipes(models.RecipeWhere.ID.EQ(int(id))).One(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	if recipe.AuthorID != authorizeUserPayload.ID.String() {
		return nil, status.Error(codes.PermissionDenied, "only the author can delete recipes")
	}

	_, err = recipe.Delete(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteRecipeResponse{}, nil
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

func validateGetRecipeRequest(req *pb.GetRecipeRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id,
			validation.Required,
			validation.Min(1),
		),
	)
}

func validateDeleteRecipeRequest(req *pb.DeleteRecipeRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id,
			validation.Required,
			validation.Min(1),
		),
	)
}

func validateListRecipesRequest(req *pb.ListRecipesRequest) error {
	return nil
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
