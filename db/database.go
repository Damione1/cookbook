package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

type Dbinstance struct {
	Db *sql.DB
}

var DB Dbinstance

func ConnectDb() error {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	runDBMigration(os.Getenv("MIGRATION_PATH"), db)

	boil.SetDB(db)

	DB = Dbinstance{
		Db: db,
	}

	return nil
}

func runDBMigration(migrationURL string, db *sql.DB) {
	fmt.Printf("Running migration...\n")

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		fmt.Printf("Failed to run migration2. \n", err)
		os.Exit(2)
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Printf("No migration to run. \n")
		} else {
			fmt.Printf("Failed to run migration 2. \n", err)
			os.Exit(2)
		}
	}

}
