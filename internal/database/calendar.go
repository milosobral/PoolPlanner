package database

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Calendar represents a calendar entity in the database
type Calendar struct {
	ID         int
	Pools      []int // List of pool IDs associated with the calendar
	UniqueHash int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// // InsertCalendar inserts a new calendar into the calendar table
// func InsertCalendar(db *sql.DB, poolIDs []int) (int, string, error) {
// 	// Generate a unique hash for the pools list
// 	hash := generateUniqueHash(poolIDs)
//
// 	query := `
// 		INSERT INTO calendars (pools, unique_hash)
// 		VALUES ($1, $2)
// 		RETURNING id;
// 	`
// 	var id int
// 	err := db.QueryRow(query, pq.Array(poolIDs), hash).Scan(&id)
// 	if err != nil {
// 		return 0, "", fmt.Errorf("error inserting calendar: %v", err)
// 	}
// 	return id, hash, nil
// }
//
// // GetCalendarByID retrieves a calendar by its ID
// func GetCalendarByID(db *sql.DB, id int) (*Calendar, error) {
// 	query := `
// 		SELECT id, pools, unique_hash, created_at, updated_at
// 		FROM calendars
// 		WHERE id = $1;
// 	`
// 	var calendar Calendar
// 	err := db.QueryRow(query, id).Scan(&calendar.ID, pq.Array(&calendar.Pools), &calendar.UniqueHash, &calendar.CreatedAt, &calendar.UpdatedAt)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, fmt.Errorf("calendar with ID %d not found", id)
// 		}
// 		return nil, fmt.Errorf("error retrieving calendar: %v", err)
// 	}
// 	return &calendar, nil
// }
//
// // GetCalendarsByPoolID retrieves calendars associated with a specific pool ID
// func GetCalendarsByPoolID(db *sql.DB, poolID int) ([]Calendar, error) {
// 	query := `
// 		SELECT id, pools, unique_hash, created_at, updated_at
// 		FROM calendars
// 		WHERE $1 = ANY(pools);
// 	`
// 	rows, err := db.Query(query, poolID)
// 	if err != nil {
// 		return nil, fmt.Errorf("error retrieving calendars by pool ID: %v", err)
// 	}
// 	defer rows.Close()
//
// 	var calendars []Calendar
// 	for rows.Next() {
// 		var calendar Calendar
// 		if err := rows.Scan(&calendar.ID, pq.Array(&calendar.Pools), &calendar.UniqueHash, &calendar.CreatedAt, &calendar.UpdatedAt); err != nil {
// 			return nil, fmt.Errorf("error scanning calendar: %v", err)
// 		}
// 		calendars = append(calendars, calendar)
// 	}
//
// 	if err = rows.Err(); err != nil {
// 		return nil, fmt.Errorf("error with rows: %v", err)
// 	}
//
// 	return calendars, nil
// }
//
// // GetAllCalendars retrieves all calendars from the calendar table
// func GetAllCalendars(db *sql.DB) ([]Calendar, error) {
// 	query := `
// 		SELECT id, pools, unique_hash, created_at, updated_at
// 		FROM calendars;
// 	`
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("error retrieving calendars: %v", err)
// 	}
// 	defer rows.Close()
//
// 	var calendars []Calendar
// 	for rows.Next() {
// 		var calendar Calendar
// 		if err := rows.Scan(&calendar.ID, pq.Array(&calendar.Pools), &calendar.UniqueHash, &calendar.CreatedAt, &calendar.UpdatedAt); err != nil {
// 			return nil, fmt.Errorf("error scanning calendar: %v", err)
// 		}
// 		calendars = append(calendars, calendar)
// 	}
//
// 	if err = rows.Err(); err != nil {
// 		return nil, fmt.Errorf("error with rows: %v", err)
// 	}
//
// 	return calendars, nil
// }

// generateUniqueHash generates a unique hash for a list of pool IDs
func GenerateUniqueHash(poolIDs []string) string {
	string := ""
	for _, id := range poolIDs {
		string += id
	}
	hash := sha256.New()
	hash.Write([]byte(string))

	hashed := hex.EncodeToString(hash.Sum(nil)[:5])
	return hashed
}
