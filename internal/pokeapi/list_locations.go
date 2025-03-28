package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c Client) ListLocations(pageUrl string, cfg *Config) (RespLocations, error) {
	url := BaseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
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
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error reading json data: %v", res.Status)
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error decoding data: %v", res.Status)
	}
	cfg.PokeCache.Add(url, bytes)

	return data, nil
}
