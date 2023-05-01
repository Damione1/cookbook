package pbx

import (
	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
)

func DbRecipeToProto(recipe *models.Recipe) *pb.Recipe {
	recipePb := &pb.Recipe{
		Id:    int64(recipe.ID),
		Title: recipe.Title,
	}
	return recipePb
}

func ProtoRecipeToDb(recipe *pb.Recipe) *models.Recipe {
	recipeDb := &models.Recipe{
		Title: recipe.GetTitle(),
	}
	if recipe.GetId() != 0 {
		recipeDb.ID = int(recipe.GetId())
	}

	return recipeDb
}
