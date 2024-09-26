package Handle

import (
	"bytes"
	"net/http"
	"text/template"
)

// // the error page of the site is being worked on.
func ErrorHandle(w http.ResponseWriter, StatusCodes int) {
	tmplet, er := template.ParseFiles("templates/Error.html")
	if er != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
	va := http.StatusText(StatusCodes)
	var buf bytes.Buffer
	er = tmplet.Execute(&buf, va)
	if er != nil {
		http.Error(w, "internal server error 500", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(StatusCodes)
	code, er := w.Write(buf.Bytes())
	if er != nil && code != http.StatusOK {
		http.Error(w, http.StatusText(code), http.StatusInternalServerError)
		return
	}
}
