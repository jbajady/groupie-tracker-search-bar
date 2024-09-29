package Handle

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	Func "GroupieTracker/Ressources"
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
	temple, err := template.ParseFiles("templates/Homepage.html")
	if err != nil {
		ErrorHandle(w, http.StatusInternalServerError)
		return
	}
	text := r.FormValue("text")
	// fmt.Println(text)
	var sl []int
	for _, v := range Func.Artists {
		fmt.Println("----------------->>11")
		creationdata := strconv.Itoa(v.CreationDate)
		if strings.Contains(v.Name, text) || strings.Contains(v.FirstAlbum, text) || strings.Contains(creationdata, text) {
			sl = append(sl, v.ID)
		}

		for i := 0; i < len(v.Location); i++ {
			fmt.Println("----------------->>222")
			if strings.Contains(v.Location[i], text) {
				fmt.Println(v.ID, v.Location[i], text)
				sl = append(sl, v.ID)
			}
		}
		for j := 0; j < len(v.Date); j++ {
			fmt.Println("----------------->>33")
			if strings.Contains(v.Location[j], text) {
				sl = append(sl, v.ID)
			}
		}
	}
	if Func.SearchArtist != nil {
		Func.SearchArtist = nil
	}
	fmt.Println("----------------->>", Func.SearchArtist)
	fmt.Println(len(sl))
	for _, ch := range Func.Artists {
		for i := 0; i < len(sl); i++ {
			if ch.ID == sl[i] {
				Func.SearchArtist = append(Func.SearchArtist, ch)
			}
			if i == len(sl)-1 {
				break
			}
		}
	}
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
