package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetComic(id int64) (*Comic, error) {
	url := fmt.Sprintf("http://xkcd.com/%d/info.0.json", id)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Could not Fetch %s due to %v", url, err)
		return &Comic{}, err
	}

	var data Comic

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&data); err != nil {
		fmt.Printf("Could not Fetch %s due to %v", url, err)
		return &Comic{}, err
	}

	defer resp.Body.Close()

	return &data, nil
}
