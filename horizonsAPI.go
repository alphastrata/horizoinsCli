package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func encodeReservedCharacters(str string) string {
	// Create a new URL-encoded value from the string
	encoded := url.QueryEscape(str)

	// Replace each reserved character with its encoded value
	encoded = strings.ReplaceAll(encoded, "%0A", "%0D%0A") // Convert newlines to %0D%0A
	encoded = strings.ReplaceAll(encoded, " ", "%20")
	encoded = strings.ReplaceAll(encoded, "#", "%23")
	encoded = strings.ReplaceAll(encoded, "$", "%24")
	encoded = strings.ReplaceAll(encoded, "&", "%26")
	encoded = strings.ReplaceAll(encoded, "+", "%2B")
	encoded = strings.ReplaceAll(encoded, ",", "%2C")
	encoded = strings.ReplaceAll(encoded, "/", "%2F")
	encoded = strings.ReplaceAll(encoded, ":", "%3A")
	encoded = strings.ReplaceAll(encoded, ";", "%3B")
	encoded = strings.ReplaceAll(encoded, "=", "%3D")
	encoded = strings.ReplaceAll(encoded, "?", "%3F")
	encoded = strings.ReplaceAll(encoded, "@", "%40")
	encoded = strings.ReplaceAll(encoded, "[", "%5B")
	encoded = strings.ReplaceAll(encoded, "]", "%5D")

	return encoded
}
func (api *HorizonsAPI) createURL() string {
	baseURL := "https://ssd.jpl.nasa.gov/api/horizons.api"

	params := url.Values{}
	params.Set("format", encodeReservedCharacters(api.Format))
	params.Set("COMMAND", encodeReservedCharacters(api.Command))
	params.Set("OBJ_DATA", encodeReservedCharacters(api.ObjData))
	params.Set("MAKE_EPHEM", encodeReservedCharacters(api.MakeEphem))
	params.Set("EPHEM_TYPE", encodeReservedCharacters(api.EphemType))
	params.Set("CENTER", encodeReservedCharacters(api.Center))
	params.Set("START_TIME", encodeReservedCharacters(api.StartTime))
	params.Set("STOP_TIME", encodeReservedCharacters(api.StopTime))
	params.Set("STEP_SIZE", encodeReservedCharacters(api.StepSize))
	params.Set("QUANTITIES", encodeReservedCharacters(api.Quantities))

	return fmt.Sprintf("%s?%s", baseURL, params.Encode())
}

func (api *HorizonsAPI) download() ([]byte, error) {
	url := api.createURL()

	fmt.Println(url)

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
