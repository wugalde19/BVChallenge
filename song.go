package main

// Structs of the program
type Song struct {
	song   string `json:"song,omitempty"`
	artist string `json:"artist,omitempty"`
	genre  string `json:"genre,omitempty"`
	length int    `json:"length,omitempty"`
}

type Songs []Song
