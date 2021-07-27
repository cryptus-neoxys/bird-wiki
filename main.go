package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// This creates the router and returns it
// Using this instantiate router in main
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// <-- Static Files -->

	// Declare static files directory
	staticFileDir := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	// wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// eg: GET /index.html
	// "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	
	// Bird Routes
	r.HandleFunc("/bird", GetBirdHandler).Methods("GET")
	r.HandleFunc("/bird", CreateBirdHandler).Methods("POST")

	return r
}

func main()  {
	// Router now instantiated with the above `newRouter` constructor
	r := newRouter()
	http.ListenAndServe(":8081", r)
}

func handler (rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}