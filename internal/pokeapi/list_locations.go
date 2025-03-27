package pokeapi

import (
	"encoding/json"
	"fmt"

	"net/http"
)

func (c Client) ListLocations(pageUrl *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, fmt.Errorf("networking problem: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error fulfilling the request: %v", err)
	}
	defer res.Body.Close()

	code := res.StatusCode
	if code > 299 {
		return RespLocations{}, fmt.Errorf("non-ok status code: %v", res.Status)
	}

	var data RespLocations
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error decoding data: %v", res.Status)
	}

	return data, nil
}