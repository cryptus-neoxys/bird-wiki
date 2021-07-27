package main

import (
	"io/ioutil"
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

func TestRouter(t *testing.T) {
	// using the same router from main to test
	r := newRouter()

	// new mock server from go's httptest lib
	mockServer := httptest.NewServer(r)

	// The mock server runs and exposes location in .URL
	// make GET to "/hello" route defined in router
	resp, err := http.Get(mockServer.URL + "/hello")

	// handle err
	if err != nil {
		t.Fatal(err)
	}

	// resp.StatusCode should be 200
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, but got %v", err)
	}

	// read resp, convert to string
	defer resp.Body.Close()
	// read body into bytes[] b
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	// converting bytes to string
	respStr := string(b)

	expected := "Hello World"

	if respStr != expected {
		t.Errorf("Received unexpected respose, got %s want %s", respStr, expected)
	}
}

func TestNonExistentRoute (t *testing.T) {
	// mostly same as above
	r := newRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL + "/hello", "", nil)

	if err  != nil {
		t.Fatal(err)
	}

	// status should be 405 (method not allowed)
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	// expecting an empty body
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respStr := string(b)
	expected := ""

	if respStr != expected {
		t.Errorf("Unexpected response body, got %s want %s", respStr, expected)
	}

}