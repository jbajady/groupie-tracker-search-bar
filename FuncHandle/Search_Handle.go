package Handle

import (
	"fmt"
	"net/http"
)

func SearchHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Search" {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		ErrorHandle(w, http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	fmt.Println(text)
	// var d []Func.Artest
	// fmt.Println(Func.Artists)
	// if strings.ContainsAny() {
	// }
}
