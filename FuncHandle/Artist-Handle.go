package Handle

import (
	"bytes"
	"fmt"
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
	fmt.Println("ok")

	IdArtest, err := strconv.Atoi(val)
	if err != nil || IdArtest < 1 || IdArtest > 52 {
		ErrorHandle(w, http.StatusNotFound)
		return
	}
	// valide := Func.GenriateData(IdArtest)
	// if !valide {
	// 	ErrorHandle(w, http.StatusInternalServerError)
	// 	return
	// }
	fmt.Println(Func.Date[IdArtest-1])
	DATA := Func.DataFinal{
		Artiste:  Func.Artists[IdArtest-1],
		Location: Func.Location[IdArtest-1],
		Date:     Func.Date[IdArtest-1],
		Relation: Func.Relation[IdArtest-1],
	}

	var buf bytes.Buffer
	err = temple.Execute(&buf, DATA)
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
