package main

import (
	"database/sql"
)

// Checks if an error exists
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Checks if an error exists when a db connection is trying to be established
func checkDBErr(db *sql.DB) bool {
	if db != nil {
		return true
	}
	return false
}
