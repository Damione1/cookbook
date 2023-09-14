package pbx

import (
	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/pkg/pb"
)

func DbRecipeToProto(recipe *models.Recipe) *pb.Recipe {
	recipePb := &pb.Recipe{
		Id:           recipe.ID,
		Title:        recipe.Title,
		Description:  recipe.Description.String,
		Instructions: recipe.Directions,
		AuthorId:     recipe.AuthorID,
	}

	//if recipe.R.RecipeIngredients != nil {
	//	recipePb.Ingredients = make([]*pb.Ingredient, len(recipe.R.RecipeIngredients))
	//	for i, ingredient := range recipe.R.RecipeIngredients {
	//		recipePb.Ingredients[i] = DbIngredientToProto(ingredient)
	//	}
	//}
	return recipePb
}

func ProtoRecipeToDb(recipe *pb.Recipe) *models.Recipe {
	recipeDb := &models.Recipe{
		Title: recipe.GetTitle(),
	}
	if recipe.GetId() != "" {
		recipeDb.ID = recipe.GetId()
	}
	return recipeDb
}

func DbIngredientToProto(ingredient *models.RecipeIngredient) *pb.Ingredient {
	ingredientPb := &pb.Ingredient{
		Id:       ingredient.IngredientID,
		Name:     ingredient.R.Ingredient.Name,
		Quantity: float32(ingredient.Quantity),
		Unit:     UnitEnumToProto(ingredient.Unit),
	}
	return ingredientPb
}

func UnitEnumToProto(unit models.IngredientUnitEnum) pb.Unit {
	switch unit {
	case models.IngredientUnitEnumGRAM:
		return pb.Unit_GRAM
	case models.IngredientUnitEnumKILOGRAM:
		return pb.Unit_KILOGRAM
	case models.IngredientUnitEnumMILLILITER:
		return pb.Unit_MILLILITER
	case models.IngredientUnitEnumLITER:
		return pb.Unit_LITER
	case models.IngredientUnitEnumTEASPOON:
		return pb.Unit_TEASPOON
	case models.IngredientUnitEnumTABLESPOON:
		return pb.Unit_TABLESPOON
	case models.IngredientUnitEnumCUP:
		return pb.Unit_CUP
	case models.IngredientUnitEnumPINCH:
		return pb.Unit_PINCH
	default:
		return pb.Unit_UNIT_UNSPECIFIED
	}
}

func UnitProtoToDb(unit pb.Unit) models.IngredientUnitEnum {
	switch unit {
	case pb.Unit_GRAM:
		return models.IngredientUnitEnumGRAM
	case pb.Unit_KILOGRAM:
		return models.IngredientUnitEnumKILOGRAM
	case pb.Unit_MILLILITER:
		return models.IngredientUnitEnumMILLILITER
	case pb.Unit_LITER:
		return models.IngredientUnitEnumLITER
	case pb.Unit_TEASPOON:
		return models.IngredientUnitEnumTEASPOON
	case pb.Unit_TABLESPOON:
		return models.IngredientUnitEnumTABLESPOON
	case pb.Unit_CUP:
		return models.IngredientUnitEnumCUP
	case pb.Unit_PINCH:
		return models.IngredientUnitEnumPINCH
	default:
		return models.IngredientUnitEnumUNIT_UNSPECIFIED
	}
}
