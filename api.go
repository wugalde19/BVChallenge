package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

/*
import (
    "encoding/json"
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"

    "database/sql"
    _ "github.com/mattn/go-sqlite3"
	"crypto/des"
	"crypto/des"
)
*/

// General variables of the program
var db *sql.DB

// Structs of the program
type Song struct {
	song   string `json:"song,omitempty"`
	artist string `json:"artist,omitempty"`
	genre  string `json:"genre,omitempty"`
	length int    `json:"length,omitempty"`
}

func generalSearch(w http.ResponseWriter, r *http.Request) {
	descriptor := pat.Param(r, "descriptor")
	fmt.Fprintf(w, "Descriptor, %s!\n", descriptor)
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE (instr(lower(Song.artist), lower('" + descriptor + "'))) OR (instr(lower(Song.song), lower('" + descriptor + "'))) OR (instr(lower(Gen.name), lower('" + descriptor + "')));"
	songs := getSongsArray(queryString)
	fmt.Println(songs)
	json.NewEncoder(w).Encode(songs)
}

func searchBySong(w http.ResponseWriter, r *http.Request) {
	song := pat.Param(r, "song")
	fmt.Fprintf(w, "Song, %s!", song)
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE instr(lower(Song.song), lower('" + song + "'));"
	songs := getSongsArray(queryString)
	fmt.Println(songs)
	json.NewEncoder(w).Encode(songs)
}

func searchByArtist(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "artist")
	fmt.Fprintf(w, "Artist, %s!", artist)
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE instr(lower(Song.artist), lower('" + artist + "'));"
	songs := getSongsArray(queryString)
	fmt.Println(songs)
	json.NewEncoder(w).Encode(songs)
}

func searchByGenre(w http.ResponseWriter, r *http.Request) {
	genre := pat.Param(r, "genre")
	fmt.Fprintf(w, "Genre, %s!", genre)
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE instr(lower(Gen.name), lower('" + genre + "'));"
	songs := getSongsArray(queryString)
	fmt.Println(songs)
	json.NewEncoder(w).Encode(songs)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/generalsearch/:descriptor"), generalSearch)
	mux.HandleFunc(pat.Get("/searchbysong/:song"), searchBySong)
	mux.HandleFunc(pat.Get("/searchbyartist/:artist"), searchByArtist)
	mux.HandleFunc(pat.Get("/searchbygenre/:genre"), searchByGenre)

	http.ListenAndServe("localhost:8000", mux)
}

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

// Returns an array with the founded songs
func getSongsArray(pQueryString string) []Song {
	var songs []Song
	db = connectToDB("sqlite3", "jrdd.db")

	if checkDBErr(db) {
		rows := doQuery(db, pQueryString)

		for rows.Next() {
			var newSong Song
			err := rows.Scan(&newSong.song, &newSong.artist, &newSong.genre, &newSong.length)
			checkErr(err)
			fmt.Println(newSong.song)
			fmt.Println(newSong.artist)
			fmt.Println(newSong.genre)
			fmt.Println(newSong.length)
			fmt.Println("--------------\n")

			songs = append(songs, newSong)
		}

		rows.Close()
	}

	db.Close()

	return songs
}

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
