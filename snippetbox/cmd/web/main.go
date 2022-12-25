package main

import (
	"log"
	"net/http"
)

// - handler
// - router(servermux)
// - server

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	// start a new web server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
