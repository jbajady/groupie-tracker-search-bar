package Func

import (
	"strconv"
	"strings"
)

func SearchOfData(text string) {
	checked := false
	if len(text) == 0 {
		SearchArtist = Artists
		return
	}
	if SearchArtist != nil {
		SearchArtist = nil
	}
	for i, Artist := range Artists {
		creationdata := strconv.Itoa(Artist.CreationDate)
		if strings.Contains(Artist.Name, text) || strings.Contains(Artist.FirstAlbum, text) || strings.Contains(creationdata, text) {
			checked = true
		}
		for _, ch := range Artist.Members {
			if strings.Contains(ch, text) {
				checked = true
			}
		}
		for _, locactin := range Relation.Index[i].Location {
			if strings.Contains(locactin, text) {
				checked = true
			}
		}
		for _, data := range Relation.Index[i].Date {
			if strings.Contains(data, text) {
				checked = true
			}
		}
		if checked {
			SearchArtist = append(SearchArtist, Artist)
			checked = false
		}

	}
}
