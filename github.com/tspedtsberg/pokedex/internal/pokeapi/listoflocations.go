package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (reLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	//If a Get request have been made within the interval time, the []byte data is stored in the cache's 
	//cacheentry.val. Unmarshalling the data is simply enough
	if val, ok := c.cache.Get(url); ok {
		locations := reLocations{}
		err := json.Unmarshal(val, &locations)
		if err != nil {
			return reLocations{}, err
		}
		return locations, nil
	}

	//Creating a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return reLocations{}, err
	}

	//Using exsisting client and making the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return reLocations{}, err
	}
	defer res.Body.Close()

	//reading the data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return reLocations{}, err
	}

	//unmarshal the data
	locations := reLocations{}
	if err := json.Unmarshal(data, &locations); err != nil {
		return reLocations{}, err
	}

	c.cache.Add(url, data)
	return locations, nil
}