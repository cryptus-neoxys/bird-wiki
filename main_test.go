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
	res, err := http.Get(mockServer.URL + "/hello")

	// handle err
	if err != nil {
		t.Fatal(err)
	}

	// res.StatusCode should be 200
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, but got %v", err)
	}

	// read res, convert to string
	defer res.Body.Close()
	// read body into bytes[] b
	b, err := ioutil.ReadAll(res.Body)

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

	res, err := http.Post(mockServer.URL + "/hello", "", nil)

	if err  != nil {
		t.Fatal(err)
	}

	// status should be 405 (method not allowed)
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("StatusCode should be 405, got %d", res.StatusCode)
	}

	// expecting an empty body
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	respStr := string(b)
	expected := ""

	if respStr != expected {
		t.Errorf("Unexpected response body, got %s want %s", respStr, expected)
	}

}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	// hit `GET /assets/` and get index.html 
	res, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	// statusCode should be 200
	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode should be 405, got %d", res.StatusCode)
	}

	// ofc can't test the entire HTML
	// so test the content-type header == "text/html; charset=utf-8"
	// thats oke, ig to know html is being served
	contentType := res.Header.Get("Content-Type")
	expectedContentType := "text/html; charset=utf-8"

	if contentType != expectedContentType {
		t.Errorf("Wrong content type, expected %s, got %s", expectedContentType, contentType)
	}
}