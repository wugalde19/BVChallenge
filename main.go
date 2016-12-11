package main

import "net/http"

func main() {
	mux := newRouter()
	http.ListenAndServe("localhost:8000", mux)
}
