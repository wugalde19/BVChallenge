package main

// Structs of the program
type Song struct {
	Song   string `json:"song,omitempty"`
	Artist string `json:"artist,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Length int    `json:"length,omitempty"`
}

type Songs []Song

//-----------------------------------------------------------
type SongByLength struct {
	Song   string `json:"song,omitempty"`
	Length int    `json:"length,omitempty"`
}

//-----------------------------------------------------------
type GenreList struct {
	Genre       string `json:"genre,omitempty"`
	NumOfSongs  int    `json:"number_of_songs,omitempty"`
	TotalLength int    `json:"total_length,omitempty"`
}

type Genres []GenreList
