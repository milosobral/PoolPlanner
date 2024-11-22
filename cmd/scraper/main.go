// Simple Hello world to get started
package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/milosobral/PoolPlanner/internal/database"
	"github.com/milosobral/PoolPlanner/internal/scraping"
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
	_, err = database.Migrate(uri)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrations done")

	// Get the pool url
	url, exists := os.LookupEnv("POOL_URL")
	if !exists {
		log.Fatal("POOL_URL not set")
	}

	scrapingTicker := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-scrapingTicker.C:
			// Scrape the pool list
			pools := scraping.GetPoolList(url)

			// Update the database
			for _, pool := range pools {
				err = database.UpdatePool(db, pool)
				if err != nil {
					log.Fatal(err)
				}
			}

			// Scrape the pool schedule for each pool
			// Update the database of events
		}
	}

}
