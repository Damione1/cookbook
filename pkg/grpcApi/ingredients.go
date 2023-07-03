package grpcApi

import (
	"context"
	"fmt"
	"strings"

	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (server *Server) InsertIngredient(ctx context.Context, req *pb.CreateIngredientRequest) (*pb.CreateIngredientResponse, error) {
	if _, err := server.authorizeUser(ctx); err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateCreateIngredientRequest(req); err != nil {
		return nil, err
	}

	ingredient := models.Ingredient{
		Name: req.Ingredient.Name,
	}

	err := ingredient.Insert(ctx, server.config.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.CreateIngredientResponse{
		Ingredient: &pb.Ingredient{
			Id:   int64(ingredient.ID),
			Name: ingredient.Name,
		},
	}, nil
}

func validateCreateIngredientRequest(req *pb.CreateIngredientRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Ingredient,
			validation.Required,
			validation.By(
				func(value interface{}) error {
					return validateIngredient(value.(*pb.Ingredient))
				},
			),
		),
	)
}

func (server *Server) ListIngredients(ctx context.Context, req *pb.ListIngredientsRequest) (*pb.ListIngredientsResponse, error) {
	err := validateListIngredientsRequest(req)
	if err != nil {
		return nil, err
	}

	pageSize := int(req.GetPageSize())
	page := 1
	if req.GetPageToken() > 0 {
		page = int(req.GetPageToken())
	}
	offset := (page - 1) * pageSize

	queryMods := []qm.QueryMod{
		qm.OrderBy("created_at desc"),
		qm.Limit(pageSize),
		qm.Offset(offset),
	}

	search := strings.TrimSpace(req.GetName())
	if search != "" {
		queryMods = append(queryMods, qm.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", search)))
	}

	dbIngredients, err := models.Ingredients(queryMods...).All(ctx, server.config.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list ingredients: %w", err)
	}

	ingredients := make([]*pb.Ingredient, len(dbIngredients))
	for i, dbIngredient := range dbIngredients {
		ingredients[i] = &pb.Ingredient{
			Id:   int64(dbIngredient.ID),
			Name: dbIngredient.Name,
		}
	}

	nextPageToken := 0
	if len(dbIngredients) == pageSize {
		nextPageToken = page + 1
	}

	return &pb.ListIngredientsResponse{
		Ingredients:   ingredients,
		NextPageToken: int32(nextPageToken),
	}, nil
}

func validateListIngredientsRequest(req *pb.ListIngredientsRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.PageSize, validation.Required, validation.Min(1), validation.Max(50)),
		validation.Field(&req.PageToken, validation.Min(0)),
		validation.Field(&req.Name, validation.RuneLength(0, 50)),
	)
}

func (server *Server) GetIngredient(ctx context.Context, req *pb.GetIngredientRequest) (*pb.GetIngredientResponse, error) {
	if err := validateGetIngredientRequest(req); err != nil {
		return nil, err
	}

	dbIngredient, err := models.FindIngredient(ctx, server.config.DB, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	return &pb.GetIngredientResponse{
		Ingredient: &pb.Ingredient{
			Id:   int64(dbIngredient.ID),
			Name: dbIngredient.Name,
		},
	}, nil

}

func validateGetIngredientRequest(req *pb.GetIngredientRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required, validation.Min(1)),
	)
}

// UpdateIngredient
func (server *Server) UpdateIngredient(ctx context.Context, req *pb.UpdateIngredientRequest) (*pb.UpdateIngredientResponse, error) {
	if _, err := server.authorizeUser(ctx); err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateUpdateIngredientRequest(req); err != nil {
		return nil, err
	}

	dbIngredient, err := models.FindIngredient(ctx, server.config.DB, int(req.Ingredient.Id))
	if err != nil {
		return nil, err
	}

	dbIngredient.Name = req.Ingredient.Name

	_, err = dbIngredient.Update(ctx, server.config.DB, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &pb.UpdateIngredientResponse{
		Ingredient: &pb.Ingredient{
			Id:   int64(dbIngredient.ID),
			Name: dbIngredient.Name,
		},
	}, nil

}

func validateUpdateIngredientRequest(req *pb.UpdateIngredientRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Ingredient,
			validation.Required,
			validation.By(
				func(value interface{}) error {
					return validateIngredient(value.(*pb.Ingredient))
				},
			),
		),
	)
}

func (server *Server) DeleteIngredient(ctx context.Context, req *pb.DeleteIngredientRequest) (*pb.DeleteIngredientResponse, error) {
	if _, err := server.authorizeUser(ctx); err != nil {
		return nil, unauthenticatedError(err)
	}

	if err := validateDeleteIngredientRequest(req); err != nil {
		return nil, err
	}

	dbIngredient, err := models.FindIngredient(ctx, server.config.DB, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	_, err = dbIngredient.Delete(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteIngredientResponse{}, nil
}

func validateDeleteIngredientRequest(req *pb.DeleteIngredientRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required, validation.Min(1)),
	)
}

func validateIngredients(ingredients []*pb.Ingredient) error {
	for _, ingredient := range ingredients {
		if err := validateIngredient(ingredient); err != nil {
			return err
		}
	}
	return nil
}

func validateIngredient(ingredient *pb.Ingredient) error {
	return validation.ValidateStruct(ingredient,
		validation.Field(&ingredient.Name,
			validation.Length(1, 100),
			//required if id is not set
			validation.By(
				func(value interface{}) error {
					if ingredient.Id == 0 {
						return validation.Required.Validate(value)
					}
					return nil
				},
			),
		),
		validation.Field(&ingredient.Quantity,
			validation.Required,
			validation.Min(1),
		),
		validation.Field(&ingredient.Unit,
			validation.Required,
		),
	)
}
