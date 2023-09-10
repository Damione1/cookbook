package main

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"

	database "github.com/Damione1/portfolio-playground/db"
	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/util"
)

type SeedData struct {
	Users             []models.User             `json:"users"`
	Ingredients       []models.Ingredient       `json:"ingredients"`
	Posts             []models.Post             `json:"posts"`
	Recipes           []models.Recipe           `json:"recipes"`
	RecipeIngredients []models.RecipeIngredient `json:"recipe_ingredients"`
}

func main() {
	ctx := context.Background()
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to load config")
	}
	log.Info().Msg("🌱 Opening the garden")

	_, err = database.ConnectDb(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to connect to database")
	}

	// Read the JSON file
	jsonFile, err := os.Open("db/seed/seed.json")
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to open JSON file")
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to read JSON data")
	}

	// Deserialize the JSON data
	var seedData SeedData
	err = json.Unmarshal(jsonData, &seedData)
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to deserialize JSON data")
	}

	// Insert the data into the database
	models.Users().DeleteAll(ctx, config.DB)
	for _, user := range seedData.Users {
		err = user.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert user")
		}
	}
	log.Info().Msg("🌱 Inserted users")

	models.Ingredients().DeleteAll(ctx, config.DB)
	for _, ingredient := range seedData.Ingredients {
		err = ingredient.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert ingredient")
		}
	}
	log.Info().Msg("🌱 Inserted ingredients")

	models.Recipes().DeleteAll(ctx, config.DB)
	for _, recipe := range seedData.Recipes {
		err = recipe.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert recipe")
		}
	}
	log.Info().Msg("🌱 Inserted recipes")

	models.Posts().DeleteAll(ctx, config.DB)
	for _, post := range seedData.Posts {
		err = post.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert post")
		}
	}
	log.Info().Msg("🌱 Inserted posts")

	models.RecipeIngredients().DeleteAll(ctx, config.DB)
	for _, recipeIngredient := range seedData.RecipeIngredients {
		err = recipeIngredient.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert recipe ingredient")
		}
	}
	log.Info().Msg("🌱 Inserted recipe ingredients")
}
