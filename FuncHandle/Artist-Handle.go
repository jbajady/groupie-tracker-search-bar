package Handle

import (
	"bytes"
	"net/http"
	"strconv"
	"text/template"

	Func "GroupieTracker/Ressources"
)

func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Artist" {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandle(w, http.StatusMethodNotAllowed)
		return
	}
	temple, err := template.ParseFiles("./templates/InformtionArtist.html")
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	val := r.FormValue("id")
	IdArtest, err := strconv.Atoi(val)
	if err != nil || IdArtest < 1 || IdArtest > 52 {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	Func.Artiste = Func.Artists[IdArtest-1]
	Func.Artiste.Location = Func.Relation.Index[IdArtest-1].Location
	Func.Artiste.Date = Func.Relation.Index[IdArtest-1].Date
	Func.Artiste.DatesLocation = Func.Relation.Index[IdArtest-1].DatesLocation
	var buf bytes.Buffer
	err = temple.Execute(&buf, Func.Artiste)
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
