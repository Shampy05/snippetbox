package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new servemux
	mux := http.NewServeMux()

	// Register the home function as the handler for the "/" URL pattern
	mux.HandleFunc("/", home)

	// Register the showSnippet function as the handler for the "/snippet" URL pattern
	mux.HandleFunc("/snippet", showSnippet)

	// Register the createSnippet function as the handler for the "/snippet/create" URL pattern
	mux.HandleFunc("/snippet/create", createSnippet)

	// Register the file server as the handler for all URL paths that start with "/static/"
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Start the server on port 4000 and log any errors
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
