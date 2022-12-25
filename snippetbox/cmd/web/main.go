package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// - handler
// - router(servermux)
// - server

// type config struct {
// 	addr      string
// 	staticDir string
// }

// var cfg config

func main() {
	addr := flag.String("addr", ":4000", "HTTP newtork address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// start a new web server
	err := srv.ListenAndServe()
	log.Fatal(err)
	errorLog.Fatal(err)
}
