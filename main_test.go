package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// The test works by send http req to server
	// result         request  method--route--body
	req, err := http.NewRequest("GET", "", nil)

	if(err != nil) {
		t.Fatal(err)
	}

	// recorder acts as the target of req, sort of as
	// a browser / http client. from Go's httptest lib
	recorder := httptest.NewRecorder()

	// Handler from main file
	hf := http.HandlerFunc(handler)

	// server req to recorder, executes the handler to test
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: received %v want %v", status, http.StatusOK)
	}

	expected := `Hello World`
	actual := recorder.Body.String()

	if actual != expected {
		t.Errorf("Handler returned unexpected body got %v want %v", actual, expected)
	}
}