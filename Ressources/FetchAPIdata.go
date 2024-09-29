package Func

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func FetchData(url string, c interface{}) error {
	data, err := http.Get(url)
	if err != nil {
		return err
	}
	defer data.Body.Close()
	dat, err := io.ReadAll(data.Body)
	if err != nil {
		return err
	}
	if strings.HasSuffix(url, "locations") {
		dat = []byte(strings.ReplaceAll(string(dat), "dates", "hd"))
	}
	er2 := json.Unmarshal(dat, &c)
	if er2 != nil {
		return er2
	}

	return nil
}
