package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Establishes a connection with a DataBase
func connectToDB(pDBDriver string, pDBName string) (db *sql.DB) {
	// It is not possible to connect to "https://s3.amazonaws.com/bv-challenge/jrdd.db"
	//db, err := sql.Open("sqlite3", "https://s3.amazonaws.com/bv-challenge/jrdd.db")
	db, err := sql.Open(pDBDriver, pDBName)
	checkErr(err)
	return db
}

// Receives a query string and return the query result
func doQuery(db *sql.DB, pQueryString string) *sql.Rows {
	// query
	rows, err := db.Query(pQueryString)
	checkErr(err)
	return rows
}
