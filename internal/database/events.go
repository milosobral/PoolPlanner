package database

import (
	"database/sql"
	"fmt"
	"time"
)

// Event represents an event in the database
type Event struct {
	ID                int
	PoolID            int
	Summary           string
	StartTime         time.Time
	EndTime           time.Time
	RecurrenceRule    string
	RecurrenceEndTime time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// InsertEvent inserts a new event into the events table
func InsertEvent(db *sql.DB, event Event) (int, error) {
	query := `
		INSERT INTO events (pool_id, summary, start_time, end_time, recurrence_rule, recurrence_end_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	var id int
	err := db.QueryRow(query, event.PoolID, event.Summary, event.StartTime, event.EndTime, event.RecurrenceRule, event.RecurrenceEndTime).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error inserting event: %v", err)
	}
	return id, nil
}

// GetEventByID retrieves an event by its ID
func GetEventByID(db *sql.DB, id int) (*Event, error) {
	query := `
		SELECT id, pool_id, summary, start_time, end_time, recurrence_rule, recurrence_end_time, created_at, updated_at 
		FROM events 
		WHERE id = $1;
	`
	var event Event
	err := db.QueryRow(query, id).Scan(&event.ID, &event.PoolID, &event.Summary, &event.StartTime, &event.EndTime, &event.RecurrenceRule, &event.RecurrenceEndTime, &event.CreatedAt, &event.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event with ID %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving event: %v", err)
	}
	return &event, nil
}

// GetEventsByPoolID retrieves events by pool ID
func GetEventsByPoolID(db *sql.DB, poolID int) ([]Event, error) {
	query := `
		SELECT id, pool_id, summary, start_time, end_time, recurrence_rule, created_at, updated_at 
		FROM events 
		WHERE pool_id = $1;
	`
	rows, err := db.Query(query, poolID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving events by pool ID: %v", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.PoolID, &event.Summary, &event.StartTime, &event.EndTime, &event.RecurrenceRule, &event.RecurrenceEndTime, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning event: %v", err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return events, nil
}

// GetAllEvents retrieves all events from the events table
func GetAllEvents(db *sql.DB) ([]Event, error) {
	query := `
		SELECT id, pool_id, summary, start_time, end_time, recurrence_rule, recurrence_end_time, created_at, updated_at 
		FROM events;
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving events: %v", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.PoolID, &event.Summary, &event.StartTime, &event.EndTime, &event.RecurrenceRule, &event.RecurrenceEndTime, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning event: %v", err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return events, nil
}

// GetEventByPoolName retrieves events based on the pool name
func GetEventByPoolName(db *sql.DB, poolName string) ([]Event, error) {
	query := `
		SELECT e.id, e.pool_id, e.summary, e.start_time, e.end_time, e.recurrence_rule, e.recurrence_end_time, e.created_at, e.updated_at
		FROM events e
		JOIN pools p ON e.pool_id = p.id
		WHERE p.name = $1;
	`
	rows, err := db.Query(query, poolName)
	if err != nil {
		return nil, fmt.Errorf("error retrieving events by pool name: %v", err)
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.PoolID, &event.Summary, &event.StartTime, &event.EndTime, &event.RecurrenceRule, &event.RecurrenceEndTime, &event.CreatedAt, &event.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning event: %v", err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	return events, nil
}
