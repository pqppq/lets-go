package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// - handler
// - router(servermux)
// - server

type config struct {
	addr      string
	staticDir string
}

var cfg config

func main() {
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	// addr := flag.String("addr", ":4000", "HTTP newtork address")
	// addr :=  os.Getenv("SNIPPETBOX_ADDR")
	flag.Parse()
	fmt.Print(cfg.staticDir)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %s", cfg.addr)
	// start a new web server
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
