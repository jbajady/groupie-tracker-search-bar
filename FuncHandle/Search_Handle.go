package Handle

import (
	"bytes"
	"net/http"
	"strings"
	"text/template"

	Func "GroupieTracker/Ressources"
)

// handles GET requests on the /Search path.
func SearchHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Search" {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandle(w, http.StatusMethodNotAllowed)
		return
	}
	temple, err := template.ParseFiles("templates/Homepage.html")
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	inputtext := r.FormValue("text")
	Func.SearchOfData(strings.ToLower(strings.TrimSpace(inputtext)))
	var buf bytes.Buffer
	err = temple.Execute(&buf, Func.SearchArtist)
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	_, er := w.Write(buf.Bytes())
	if er != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
}
