package main

import (
	"log"
	"net/http"
)

// define a home handler function, which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a snippet"))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home) // restrict this route to exact matches on / only.
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting on server :4000")

	//  Use the http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
