package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"goji.io/pat"
)

// General variables of the program
var db *sql.DB

// Returns a Json with the resultant songs of the songs search by genre, artist or song
func generalSearch(w http.ResponseWriter, r *http.Request) {
	descriptor := pat.Param(r, "descriptor")
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE (instr(lower(Song.artist), lower('" + descriptor + "'))) OR (instr(lower(Song.song), lower('" + descriptor + "'))) OR (instr(lower(Gen.name), lower('" + descriptor + "')));"
	songs := getSongsArray(queryString)
	json.NewEncoder(w).Encode(songs)
}

// Returns a Json with the resultant songs of the songs search by song
func searchBySong(w http.ResponseWriter, r *http.Request) {
	song := pat.Param(r, "song")
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE instr(lower(Song.song), lower('" + song + "'));"
	songs := getSongsArray(queryString)
	json.NewEncoder(w).Encode(songs)
}

// Returns a Json with the resultant songs of the songs search by artist
func searchByArtist(w http.ResponseWriter, r *http.Request) {
	artist := pat.Param(r, "artist")
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE instr(lower(Song.artist), lower('" + artist + "'));"
	songs := getSongsArray(queryString)
	json.NewEncoder(w).Encode(songs)
}

// Returns a Json with the resultant songs of the songs search by genre
func searchByGenre(w http.ResponseWriter, r *http.Request) {
	genre := pat.Param(r, "genre")
	var queryString = "SELECT Song.song AS song, Song.artist AS artist, Gen.name AS 'genre name', Song.length AS length FROM songs AS Song INNER JOIN genres AS Gen ON Gen.ID = Song.ID WHERE instr(lower(Gen.name), lower('" + genre + "'));"
	songs := getSongsArray(queryString)
	json.NewEncoder(w).Encode(songs)
}

// Returns an array with the founded songs
func getSongsArray(pQueryString string) []Song {
	var songs Songs

	db = connectToDB("sqlite3", "jrdd.db")

	if checkDBErr(db) {
		rows := doQuery(db, pQueryString)

		var vSong string
		var vArtist string
		var vGenre string
		var vLength int

		for rows.Next() {
			err := rows.Scan(&vSong, &vArtist, &vGenre, &vLength)
			checkErr(err)
			fmt.Println(vSong)
			fmt.Println(vArtist)
			fmt.Println(vGenre)
			fmt.Println(vLength)
			fmt.Println("--------------\n")

			newSong := Song{Song: vSong, Artist: vArtist, Genre: vGenre, Length: vLength}
			songs = append(songs, newSong)
		}
		rows.Close()
	}
	db.Close()
	return songs
}
