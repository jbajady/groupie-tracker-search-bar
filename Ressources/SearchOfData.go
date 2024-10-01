package Func

import (
	"strconv"
	"strings"
)

// Func  For Get  The  Artists Needs Using  Search
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
		if strings.Contains(strings.ToLower(Artist.Name), text) || strings.Contains(strings.ToLower(Artist.FirstAlbum), text) || strings.Contains(creationdata, text) {
			checked = true
		}
		for _, ch := range Artist.Members {
			if strings.Contains(strings.ToLower(ch), text) {
				checked = true
			}
		}
		for _, locactin := range Relations.Index[i].Location {
			if strings.Contains(strings.ToLower(locactin), text) {
				checked = true
			}
		}
		for _, data := range Relations.Index[i].Date {
			if strings.Contains(strings.ToLower(data), text) {
				checked = true
			}
		}
		if checked {
			SearchArtist = append(SearchArtist, Artist)
			checked = false
		}

	}
}
