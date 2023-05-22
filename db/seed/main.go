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
	Recipes           []models.Recipe           `json:"recipes"`
	RecipeIngredients []models.RecipeIngredient `json:"recipe_ingredients"`
	Post              []models.Post             `json:"posts"`
}

func main() {
	ctx := context.Background()
	log.Info().Msg("🌱 Seeding database")
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("👋 Failed to load config")
	}

	db, err := database.ConnectDb(&config)
	if err != nil {
		log.Fatal().Err(err).Msg("👋 Failed to connect to database")
	}

	// Read the JSON file
	jsonFile, err := os.Open("seed_data.json")
	if err != nil {
		log.Fatal().Err(err).Msg("Error opening JSON file")
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading JSON data")
	}

	// Deserialize the JSON data
	var seedData SeedData
	err = json.Unmarshal(jsonData, &seedData)
	if err != nil {
		log.Fatal().Err(err).Msg("Error unmarshalling JSON data")
	}

	for _, user := range seedData.Users {
		err := user.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating user")
		}
	}

	for _, ingredient := range seedData.Ingredients {
		err := ingredient.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating ingredient")
		}
	}

	for _, recipe := range seedData.Recipes {
		err := recipe.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating recipe")
		}
	}

	for _, recipeIngredient := range seedData.RecipeIngredients {
		err := recipeIngredient.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating recipe ingredient")
		}
	}

	for _, post := range seedData.Post {
		err := post.Insert(ctx, db, boil.Infer())
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating post")
		}
	}

	log.Info().Msg("🌱 Seeding complete")
}
