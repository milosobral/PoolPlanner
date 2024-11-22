package database

import (
	"database/sql"
	"fmt"
)

// Pool represents a pool entity in the database
type Pool struct {
	ID           int
	Name         string
	Href         string
	Address      string
	Neighborhood string
	CreatedAt    string
	UpdatedAt    string
}

// Method to initialize a Pool object from the scraping data
func PoolFromScrapingData(name, href, address, neighborhood string) Pool {
	return Pool{
		Name:         name,
		Href:         href,
		Address:      address,
		Neighborhood: neighborhood,
	}
}

func UpdatePool(db *sql.DB, pool Pool) error {
	// Check if the pool is already in the database
	exists, err := PoolExists(db, pool.Name)
	if err != nil {
		return fmt.Errorf("error checking pool: %v", err)
	}
	if !exists {
		// If the pool is not in the database, insert it
		_, err := InsertPool(db, pool)
		if err != nil {
			return fmt.Errorf("error inserting pool: %v", err)
		}
	} else {
		// If the pool is already in the database, update it
		err := UpdatePoolByID(db, pool)
		if err != nil {
			return fmt.Errorf("error updating pool: %v", err)
		}
	}
	return nil
}

func UpdatePoolByID(db *sql.DB, pool Pool) error {
	// Update the Pools values in the database except for the name and the created at
	query := `
		UPDATE pools 
		SET href = $1, address = $2, neighborhood = $3, updated_at = NOW() 
		WHERE name = $4;
	`
	_, err := db.Exec(query, pool.Href, pool.Address, pool.Neighborhood, pool.Name)
	if err != nil {
		return err
	}
	return nil
}

// Check if a pool exists in the database based on name
func PoolExists(db *sql.DB, name string) (bool, error) {
	query := `
		SELECT id 
		FROM pools 
		WHERE name = $1;
	`
	var id int
	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error checking pool: %v", err)
	}
	return true, nil
}

// InsertPool inserts a new pool into the pools table
func InsertPool(db *sql.DB, pool Pool) (int, error) {
	query := `
		INSERT INTO pools (name, href, address, neighborhood) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id;
	`
	var id int
	err := db.QueryRow(query, pool.Name, pool.Href, pool.Address, pool.Neighborhood).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting pool: %v", err)
	}
	return id, nil
}
