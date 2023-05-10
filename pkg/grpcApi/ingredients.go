package grpcApi

import (
	"context"
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
	pageToken := int(req.GetPageToken())

	// Set default page size if not provided or if it's greater than the maximum allowed
	var maxPageSize int = 50
	if pageSize <= 0 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	offset := pageSize * pageToken

	queryMods := []qm.QueryMod{
		qm.OrderBy("created_at desc"),
		qm.Limit(pageSize),
		qm.Offset(offset),
	}

	if search := strings.TrimSpace(req.GetName()); search != "" {
		queryMods = append(queryMods, qm.Where("LOWER(name) ILIKE ?", strings.ToLower("%"+search+"%")))
	}

	dbIngredients, err := models.Ingredients(queryMods...).All(ctx, server.config.DB)
	if err != nil {
		return nil, err
	}

	posts := make([]*pb.Ingredient, 0, len(dbIngredients))
	for _, dbIngredient := range dbIngredients {
		posts = append(posts,
			&pb.Ingredient{
				Id:   int64(dbIngredient.ID),
				Name: dbIngredient.Name,
			},
		)
	}

	nextPageToken := 0
	if len(dbIngredients) == pageSize {
		nextPageToken = pageToken + 1
	}

	return &pb.ListIngredientsResponse{
		Ingredients:   posts,
		NextPageToken: int32(nextPageToken),
	}, nil
}

func validateListIngredientsRequest(req *pb.ListIngredientsRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.PageSize, validation.Required, validation.Min(1), validation.Max(50)),
		validation.Field(&req.PageToken, validation.Min(0)),
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

func validateIngredient(ingredient *pb.Ingredient) error {
	return validation.ValidateStruct(ingredient,
		validation.Field(&ingredient.Name,
			validation.Required,
			validation.Length(1, 100),
		),
	)
}
