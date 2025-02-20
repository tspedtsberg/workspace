package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName
	
	//If a Get request have been made within the interval time, the []byte data is stored in the cache's 
	//cacheentry.val. Unmarshalling the data is simply enough
	if val, ok := c.cache.Get(url); ok {
		location := Location{}
		err := json.Unmarshal(val, &location)
		if err != nil {
			return Location{}, err
		}
		return location, nil
	}

	//Creating the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	//using exsisting client and making the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}

	//closing the request
	defer res.Body.Close()

	//reading the data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	//decoding the data
	location := Location{}
	if err := json.Unmarshal(data, &location); err != nil {
		return Location{}, err
	}

	//adding to cache
	c.cache.Add(url, data)
	return location, nil
}