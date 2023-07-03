package database

import (
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/Damione1/portfolio-playground/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ConnectDb(config *util.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s TimeZone=America/New_York sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDb,
	))
	if err != nil {
		log.Fatal().Err(err).Msg("🥝 Failed to connect to database")
	}

	boil.SetDB(db)

	config.DB = db

	return db, nil
}

func RunDBMigration(migrationURL string, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("🥝 Failed to create migration driver")
	}

	m, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		log.Fatal().Err(err).Msg("🥝 Failed to create migration instance")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Error().Err(err).Msg("🥝 Migration failed, attempting to recover")

		if mErr := m.Force(-1); mErr != nil {
			log.Fatal().Err(mErr).Msg("🥝 Failed to revert failed migration")
		} else {
			log.Print("🥝 Reverted failed migration")

			if err = m.Up(); err != nil {
				log.Fatal().Err(err).Msg("🥝 Failed to rerun migration after revert")
			}
		}
	} else {
		log.Print("🥝 No migration to run")
	}
	return
}
