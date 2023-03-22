package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Damione1/portfolio-playground/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

func ConnectDb(config *util.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s TimeZone=America/New_York sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDb,
	))
	if err != nil {
		fmt.Printf("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	runDBMigration(config.MigrationPath, db)

	boil.SetDB(db)

	config.DB = db

	return db, nil
}

func runDBMigration(migrationURL string, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		fmt.Printf("Failed to create database driver. %v \n", err)
		os.Exit(2)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		fmt.Printf("Failed to create migration instance. %v \n", err)
		os.Exit(2)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Printf("No migration to run. \n")
		} else {
			fmt.Printf("Failed to run migration. \n", err)
			os.Exit(2)
		}
	}
	fmt.Printf("Migration successful. \n")

	return
}
