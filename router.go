package main

import (
	"goji.io"
	"goji.io/pat"
)

// Returns a mux handler
func newRouter() *goji.Mux {
	mux := goji.NewMux()

	for _, route := range routes {
		assignRoute(route, mux)
	}

	return mux
}

// Assigns routes to the mux handler
func assignRoute(pRoute Route, pMux *goji.Mux) {
	if pRoute.Method == "GET" {
		pMux.HandleFunc(pat.Get(pRoute.Pattern), pRoute.HandlerFunc)
	}
}
