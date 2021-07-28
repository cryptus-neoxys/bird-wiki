package main

import "database/sql"

// Store has 2 Methods, each returning an err
// Get: to return all birds in DB
// Create: takes a bird as input and stores in DB
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}

// dbStore will implement store interface
// also takes dbCon obj, representing db connection
type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBird(bird *Bird) error {
	// 'Bird' has "species" and "description" attributes
	// The first underscore means data isn't required from 
	// insert query. if it was inserted correctly,
	// and the error will be populated if it wasn't
	_, err := store.db.Query("INSERT INTO birds(species, description) VALUES(?, ?)", bird.Species, bird.Description)
	return err
}

func (store *dbStore) GetBirds() ([] *Bird, error) {
	// Query the DB for all birds and store the result
	// into rows
	rows, err := store.db.Query("SELECT species, description from birds")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create birds array to be returned
	birds := [] *Bird{}

	for rows.Next() {
		// for each row, create a pointer to a bird
		bird := &Bird{}
		// Populate attributes and return err in case any
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		// If all goes well append bird to returning list
		birds = append(birds, bird)
	}

	return birds, nil
}

// Package Level Variable, available throughout
var  store Store

// Used to initialise store at start of application
// i.e. when starting server / mocking API
func initStore (s Store) {
	store = s
}