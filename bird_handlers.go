package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func GetBirdHandler (rw http.ResponseWriter, r *http.Request) {
	birdListBytes, err := json.Marshal(birds)

	// if err -> print and return
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if no err -> return list of birds
	rw.Write(birdListBytes)
}

func CreateBirdHandler (rw http.ResponseWriter, r *http.Request) {
	// new bird instance
	bird := Bird{}

	// using html forms to send data
	// `parseForm` method parses form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the values from form and create bird
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	// add bird to the list
	birds = append(birds, bird)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(rw, r, "/assets/", http.StatusFound)
}