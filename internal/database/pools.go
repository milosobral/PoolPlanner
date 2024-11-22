package database

import (
	"database/sql"
	"fmt"
)

// Pool represents a pool entity in the database
type Pool struct {
	ID        int
	Name      string
	Location  string
	CreatedAt string
	UpdatedAt string
}

// InsertPool inserts a new pool into the pools table
func InsertPool(db *sql.DB, name, location string) (int, error) {
	query := `
		INSERT INTO pools (name, location) 
		VALUES ($1, $2) 
		RETURNING id;
	`
	var id int
	err := db.QueryRow(query, name, location).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting pool: %v", err)
	}
	return id, nil
}

// GetPoolByID retrieves a pool by its ID
func GetPoolByID(db *sql.DB, id int) (*Pool, error) {
	query := `
		SELECT id, name, location, created_at, updated_at 
		FROM pools 
		WHERE id = $1;
	`
	var pool Pool
	err := db.QueryRow(query, id).Scan(&pool.ID, &pool.Name, &pool.Location, &pool.CreatedAt, &pool.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("pool with ID %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving pool: %v", err)
	}
	return &pool, nil
}

// GetPoolsByNames retrieves IDs for multiple pools by their names
func GetPoolsByNames(db *sql.DB, name string) ([]int, error) {
	query := `
		SELECT id 
		FROM pools 
		WHERE name = $1;
	`
	rows, err := db.Query(query, name)
	if err != nil {
		return nil, fmt.Errorf("error retrieving pool IDs: %v", err)
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, fmt.Errorf("error scanning pool ID: %v", err)
		}
		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return ids, nil
}

// GetAllPools retrieves all pools from the pools table
func GetAllPools(db *sql.DB) ([]Pool, error) {
	query := `
		SELECT id, name, location, created_at, updated_at 
		FROM pools;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving pools: %v", err)
	}
	defer rows.Close()

	var pools []Pool
	for rows.Next() {
		var pool Pool
		if err := rows.Scan(&pool.ID, &pool.Name, &pool.Location, &pool.CreatedAt, &pool.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning pool: %v", err)
		}
		pools = append(pools, pool)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return pools, nil
}
