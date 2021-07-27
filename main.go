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