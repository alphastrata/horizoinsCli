package main

import (
	"fmt"
	"testing"
)

func TestHorizonsAPI_download(t *testing.T) {
	// Create a HorizonsAPI object with the test server's URL
	api := &HorizonsAPI{
		Format:     "text",
		Command:    "499",
		ObjData:    "YES",
		MakeEphem:  "NO",
		EphemType:  "OBSERVER",
		Center:     "500@399",
		StartTime:  "2006-01-01",
		StopTime:   "2006-01-20",
		StepSize:   "1%20d",
		Quantities: "QUANTITIES='29'",
	}

	// Make the API call and check that the response matches the expected response
	actualResponse, err := api.download()
	if err != nil {
		t.Errorf("Error downloading data: %s", err)
	}
	fmt.Println(string(actualResponse))
}

// format=text&COMMAND='499'&OBJ_DATA='YES'&MAKE_EPHEM='YES'&EPHEM_TYPE='OBSERVER'&CENTER='500@399'&START_TIME='2006-01-01'&STOP_TIME='2006-01-20'&STEP_SIZE='1%20d'&QUANTITIES='1,9,20,23,24,29'
