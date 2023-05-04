package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HorizonsAPI struct {
	Format     string
	Command    string
	ObjData    string
	MakeEphem  string
	EphemType  string
	Center     string
	StartTime  string
	StopTime   string
	StepSize   string
	Quantities string
}

func (api *HorizonsAPI) createURL() string {
	baseURL := "https://ssd.jpl.nasa.gov/api/horizons.api"
	params := url.Values{}
	params.Set("format", api.Format)
	params.Set("COMMAND", api.Command)
	params.Set("OBJ_DATA", api.ObjData)
	params.Set("MAKE_EPHEM", api.MakeEphem)
	params.Set("EPHEM_TYPE", api.EphemType)
	params.Set("CENTER", api.Center)
	params.Set("START_TIME", api.StartTime)
	params.Set("STOP_TIME", api.StopTime)
	params.Set("STEP_SIZE", api.StepSize)
	params.Set("QUANTITIES", api.Quantities)
	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func (api *HorizonsAPI) download() ([]byte, error) {
	url := api.createURL()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
