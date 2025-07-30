package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Use the Header().Add() method to add a 'Server: Go' header to the
	// response header map. The first parameter is the header name, and
	// the second parameter is the header value.
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying a sp ecific snippet with ID %d...", id)
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// send a 201 status code.
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("save a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	// Prefix the route patterns with the required HTTP method (for now, we will
	// restrict all three routes to acting on GET requests).
	mux.HandleFunc("GET /{$}", home) // restrict this route to exact matches on / only.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting on server :4000")

	//  Use the http.ListenAndServe() function to start a new web server.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
