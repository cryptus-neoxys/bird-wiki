package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)



func TestGetBirdsHandler(t *testing.T) {
	// init mock store
	mockStore := InitMockStore()

	// Define data we want to be returned upon calling `GetBirds`
	// Also expect it to be called once

	mockStore.On("GetBirds").Return([]*Bird {
		{"sparrow", "A smol bird"},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(GetBirdHandler)

	// Now calling the handler hits mock store instead of DB
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Bird{"sparrow", "A smol bird"}
	b := []Bird{}

	err = json.NewDecoder(recorder.Body).Decode(&b)
	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body, got %v want %v", actual, expected)
	}

	mockStore.AssertExpectations(t)
}


func TestCreateBirdsHandler(t *testing.T) {

	mockStore := InitMockStore()
	/*
	 Similarly, we define our expectations for th `CreateBird` method.
	 We expect the first argument to the method to be the bird struct
	 defined below, and tell the mock to return a `nil` error
	*/
	mockStore.On("CreateBird", &Bird{"eagle", "A bird of prey"}).Return(nil)

	form := newCreateBirdForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(CreateBirdHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func newCreateBirdForm() *url.Values {
	form := url.Values{}
	form.Set("species", "eagle")
	form.Set("description", "A bird of prey")
	return &form
}