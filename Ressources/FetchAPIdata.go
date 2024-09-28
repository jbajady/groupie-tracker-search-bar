package Func

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	er2 := json.Unmarshal(dat, &c)
	fmt.Println("fetche data--------------------->", c)
	if er2 != nil {
		return er2
	}

	return nil
}
