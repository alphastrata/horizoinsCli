package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHorizonsAPI_download(t *testing.T) {
	// Create a test server that always returns the same response
	expectedResponse := []byte("Test response")
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expectedResponse)
	}))
	defer testServer.Close()

	// Create a HorizonsAPI object with the test server's URL
	api := &HorizonsAPI{
		Format:     "text",
		Command:    "'499'",
		ObjData:    "'YES'",
		MakeEphem:  "'YES'",
		EphemType:  "'OBSERVER'",
		Center:     "'500@399'",
		StartTime:  "'2006-01-01'",
		StopTime:   "'2006-01-20'",
		StepSize:   "'1 d'",
		Quantities: "'1,9,20,23,24,29'",
	}

	// Make the API call and check that the response matches the expected response
	actualResponse, err := api.download()
	if err != nil {
		t.Errorf("Error downloading data: %s", err)
	}
	fmt.Println(actualResponse)
}
