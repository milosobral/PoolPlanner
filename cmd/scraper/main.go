// Simple Hello world to get started
package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/milosobral/PoolPlanner/internal/database"
)

func main() {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Check if we are in DEV or PROD Mode
	if os.Getenv("ENV") != "DEV" {
		log.Fatal("Not in DEV mode")
	}

	// Create the database connection object
	uri, exists := os.LookupEnv("DB_URI_DEV")
	if !exists {
		log.Fatal("DB_URI_DEV not set")
	}
	db, err := database.Connect(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Actually establish the connection to the database
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to the database")

	// Do the migrations
	m, err := database.Migrate(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrations done")

	// Restore the Database
	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
	log.Println("Database restored")

}
