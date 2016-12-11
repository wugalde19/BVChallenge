package main

// Structs of the program
type Song struct {
	Song   string `json:"song,omitempty"`
	Artist string `json:"artist,omitempty"`
	Genre  string `json:"genre,omitempty"`
	Length int    `json:"length,omitempty"`
}

type Songs []Song
