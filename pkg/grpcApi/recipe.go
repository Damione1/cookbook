package grpcApi

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	"github.com/Damione1/portfolio-playground/pkg/pbx"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
		AuthorID:    authorizeUserPayload.UserID,
	}

	for _, ingredient := range req.Recipe.Ingredients {
		recipeIngredient := &models.RecipeIngredient{
			RecipeID:     recipe.ID,
			Quantity:     float64(ingredient.Quantity),
			IngredientID: int(ingredient.Id),
		}
		if ingredient.Id == 0 {
			recipeIngredient.ManualName = null.StringFrom(ingredient.Name)
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
	recipe, err := models.Recipes(
		qm.Load(models.RecipeRels.RecipeIngredients),
		qm.Load(qm.Rels(models.RecipeRels.RecipeIngredients, models.RecipeIngredientRels.Ingredient)),
		models.RecipeWhere.ID.EQ(int(req.Id)),
	).One(ctx, server.config.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "recipe not found")
		}
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

	pageSize := int(req.GetPageSize())
	pageToken := 1
	if req.GetPageToken() != 0 {
		pageToken = int(req.GetPageToken())
	}
	offset := (pageToken - 1) * pageSize

	dbRecipes, err := models.Recipes(
		qm.OrderBy("created_at desc"),
		qm.Limit(pageSize),
		qm.Offset(offset),
	).All(ctx, server.config.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &pb.ListRecipesResponse{}, nil
		}
		return nil, err
	}

	recipes := make([]*pb.Recipe, len(dbRecipes))
	for i, recipe := range dbRecipes {
		recipes[i] = &pb.Recipe{
			Id:          int64(recipe.ID),
			Title:       recipe.Title,
			Description: recipe.Description.String,
			AuthorId:    recipe.AuthorID,
		}
	}

	nextPageToken := 0
	if len(recipes) == pageSize {
		nextPageToken = pageToken + 1
	}

	return &pb.ListRecipesResponse{
		Recipes:       recipes,
		NextPageToken: int32(nextPageToken),
	}, nil
}

func (server *Server) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*pb.DeleteRecipeResponse, error) {
	authorizeUserPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateDeleteRecipeRequest(req); err != nil {
		return nil, err
	}

	recipe, err := models.Recipes(models.RecipeWhere.ID.EQ(int(req.Id))).One(ctx, server.config.DB)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "recipe not found")
		}
		return nil, err
	}

	if recipe.AuthorID != authorizeUserPayload.UserID {
		return nil, status.Error(codes.PermissionDenied, "nice try, but you can't delete someone else's recipe")
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
	return validation.ValidateStruct(req,
		validation.Field(&req.PageSize,
			validation.Required,
			validation.Min(1),
			validation.Max(100),
		),
		validation.Field(&req.PageToken,
			validation.Min(1),
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
