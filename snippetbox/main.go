package main

import (
	"log"
	"net/http"
)

// - handler
// - router(servermux)
// - server

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Sinppetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a specifiec snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowd"))
		return
	}
	w.Write([]byte("Creawte a new snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on :4000")
	// start a new web server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
