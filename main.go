package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// define a home handler function, which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
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
