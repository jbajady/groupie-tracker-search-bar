package Func

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchData(url string, c interface{}) error {
	data, err := http.Get(url)
	if err != nil {
		return err
	}
	dat, err := io.ReadAll(data.Body)
	if err != nil {
		return err
	}
	er2 := json.Unmarshal(dat, &c)
	if er2 != nil {
		return er2
	}

	return nil
}
