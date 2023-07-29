package seed

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"

	database "github.com/Damione1/portfolio-playground/db"
	"github.com/Damione1/portfolio-playground/db/models"
	"github.com/Damione1/portfolio-playground/util"
	// Import your database connection and ORM packages
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

	_, err = database.ConnectDb(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to connect to database")
	}

	// Read the JSON file
	jsonFile, err := os.Open("seed_data.json")
	if err != nil {
		log.Fatal().Err(err).Msg("🌱 Failed to open JSON file")
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
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
	for _, user := range seedData.Users {
		err = user.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert user")
		}
	}

	for _, ingredient := range seedData.Ingredients {
		err = ingredient.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert ingredient")
		}
	}

	for _, recipe := range seedData.Recipes {
		err = recipe.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert recipe")
		}
	}

	for _, post := range seedData.Posts {
		err = post.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert post")
		}
	}

	for _, recipeIngredient := range seedData.RecipeIngredients {
		err = recipeIngredient.Insert(ctx, config.DB, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("🌱 Failed to insert recipe ingredient")
		}
	}
}
