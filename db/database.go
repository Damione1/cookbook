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

func ConnectDb(config util.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s sslmode=disable TimeZone=America/New_York",
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDb,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	fmt.Printf("Running database migration %s", config.MigrationPath)

	runDBMigration(config.MigrationPath, db)

	boil.SetDB(db)

	return db, nil
}

func runDBMigration(migrationURL string, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		fmt.Printf("Failed to create migration instance. %v \n", err)
		os.Exit(2)
	}

	fmt.Printf("Running migration")
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Printf("No migration to run. \n")
		} else {
			fmt.Printf("Failed to run migration. \n", err)
			os.Exit(2)
		}
	}

}
