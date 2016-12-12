package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"General Search",
		"GET",
		"/generalsearch/:descriptor",
		generalSearch,
	},
	Route{
		"Search By Song",
		"GET",
		"/searchbysong/:song",
		searchBySong,
	},
	Route{
		"Seacrh By Artist",
		"GET",
		"/searchbyartist/:artist",
		searchByArtist,
	},
	Route{
		"Seacrh By Genre",
		"GET",
		"/searchbygenre/:genre",
		searchByGenre,
	},
	Route{
		"List of Genres",
		"GET",
		"/genreslist",
		getGenresList,
	},
}
