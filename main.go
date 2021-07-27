package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")

	http.ListenAndServe(":8081", r)
}

func handler (rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World")
}