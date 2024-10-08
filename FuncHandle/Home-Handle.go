package Handle

import (
	"bytes"
	"net/http"
	"text/template"

	Func "GroupieTracker/Ressources"
)

// the function only handles GET requests on the root path (/)
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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
	errr := Func.FetchData("https://groupietrackers.herokuapp.com/api", &Func.API)
	if errr != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	if !Func.GenriateData() {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}

	if len(Func.Artists) == 0 {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	var buf bytes.Buffer
	err = temple.Execute(&buf, Func.Artists)
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
