package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	//If a Get request have been made within the interval time, the []byte data is stored in the cache's 
	//cacheentry.val. Unmarshalling the data is simply enough
	if val, ok := c.cache.Get(url); ok {
		PokeM := Pokemon{}
		if err := json.Unmarshal(val, &PokeM); err != nil {
			return Pokemon{}, err
		}
		return PokeM, nil
	}

	//Creating request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	//using exsisten client and making the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	//closing request
	defer res.Body.Close()

	//reading the data
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	//decoding the data
	PokeM := Pokemon{}
	if err := json.Unmarshal(data, &PokeM); err != nil {
		return Pokemon{}, err
	}

	//add to cache
	c.cache.Add(url, data)
	return PokeM, nil
}