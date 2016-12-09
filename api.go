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
)
*/

type Song struct {
	song   string `json:"song,omitempty"`
	artist string `json:"artist,omitempty"`
	genre  string `json:"genre,omitempty"`
	length int    `json:"length,omitempty"`
}

func getSongs() {
	Song.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func generalSearch(w http.ResponseWriter, r *http.Request) {
	descriptor := pat.Param(r, "descriptor")
	fmt.Fprintf(w, "Descriptor, %s!", descriptor)
}

func searchBySong(w http.ResponseWriter, r *http.Request) {
	song := pat.Param(r, "song")
	fmt.Fprintf(w, "Song, %s!", song)
}

func searchByArtist(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "artist")
	fmt.Fprintf(w, "Artist, %s!", artist)
}

func searchByGenre(w http.ResponseWriter, r *http.Request) {
	genre := pat.Param(r, "genre")
	fmt.Fprintf(w, "Genre, %s!", genre)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/generalsearch/:descriptor"), generalSearch)
	mux.HandleFunc(pat.Get("/searchbysong/:song"), searchBySong)
	mux.HandleFunc(pat.Get("/searchbyartist/:artist"), searchByArtist)
	mux.HandleFunc(pat.Get("/searchbygenre/:genre"), searchByGenre)

	http.ListenAndServe("localhost:8000", mux)
}

func connectToDB() {
	//db, err := sql.Open("sqlite3", "https://s3.amazonaws.com/bv-challenge/jrdd.db")
	db, err := sql.Open("sqlite3", "jrdd.db")
	checkErr(err)

	// query
	rows, err := db.Query("SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE (instr(lower(Song.artist), lower('rock'))) OR (instr(lower(Song.song), lower('rock'))) OR (instr(lower(Gen.name), lower('rock')));")
	checkErr(err)

	var song string
	var artist string
	var genre string
	var length int

	for rows.Next() {
		err = rows.Scan(&song, &artist, &genre, &length)
		checkErr(err)
		fmt.Println(song)
		fmt.Println(artist)
		fmt.Println(genre)
		fmt.Println(length)
	}

	rows.Close() //good habit to close

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
