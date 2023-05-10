package seed

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"

	database "github.com/Damione1/portfolio-playground/db"
	"github.com/Damione1/portfolio-playground/util"
	// Import your database connection and ORM packages
)

type SeedData struct {
	Users       []User       `json:"users"`
	Ingredients []Ingredient `json:"ingredients"`
	Recipes     []Recipe     `json:"recipes"`
	Post        []Post       `json:"posts"`
}

func main() {
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
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Deserialize the JSON data
	var seedData SeedData
	err = json.Unmarshal(jsonData, &seedData)
	if err != nil {
		log.Fatalf("Error deserializing JSON data: %v", err)
	}

	// Insert the data into the database
	// Use your ORM or raw SQL queries to insert the data
	// ...
}
