// Simple Hello world to get started
package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/milosobral/PoolPlanner/internal/database"
	"github.com/milosobral/PoolPlanner/internal/scraping"
)

// Global var for the log level
var logLevel slog.LevelVar

func main() {
	// Set the Logging level
	logLevel.Set(slog.LevelDebug)
	log := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: &logLevel, // Use LevelVar
	}))

	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading environment variables: %v", "Error", err)
	}

	// Check if we are in DEV or PROD Mode
	if os.Getenv("ENV") != "DEV" {
		log.Error("Not in DEV mode")
	}

	// Create the database connection object
	uri, exists := os.LookupEnv("DB_URI_DEV")
	if !exists {
		log.Error("DB_URI_DEV not set")
	}
	db, err := database.Connect(uri)
	if err != nil {
		log.Error("Error creating the database connection object: %v", "Error", err)
	}
	defer db.Close()

	// Actually establish the connection to the database
	err = db.Ping()
	if err != nil {
		log.Error("Error establishing the connection to the database: %v", "Error", err)
	}
	log.Info("Connected to the database")

	// Do the migrations
	_, err = database.Migrate(uri)
	if err != nil {
		log.Error("Error migrating the database: %v", "Error", err)
	}
	log.Info("Migrations done")

	// Get the pool url
	url, exists := os.LookupEnv("POOL_URL")
	if !exists {
		log.Error("POOL_URL not set")
	}

	log.Info("Starting the scraper...")

	// Interval for the scraping
	interval := time.Minute * 5
	scrapingTicker := time.NewTicker(interval)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

loop:
	for {

		// Scrape the pool list
		log.Debug("Scraping the pool list")
		pools := scraping.GetPoolList(url)
		log.Debug(fmt.Sprintf("Found %d pools", len(pools)))

		// Update the database
		log.Debug("Updating the database")
		for _, pool := range pools {
			err = database.UpdatePool(db, pool)
			if err != nil {
				log.Error(fmt.Sprintf("Error updating pool: %v", pool.Name), "Error", err)
			}
		}
		log.Debug("Database updated")

		// Scrape the pool schedule for each pool
		// TODO: Add the code to update the schedules

		// Update the database of events
		// TODO: Add the code to update the events

		log.Debug(fmt.Sprintf("Waiting for %s", interval))

		// Loop to wait for the next tick
		select {
		case <-scrapingTicker.C:
			continue
		case <-interrupt:
			scrapingTicker.Stop()
			break loop
		}
	}

}
