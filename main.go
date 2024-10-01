package main

import (
	"fmt"
	"net/http"

	Handle "GroupieTracker/FuncHandle"
)

func main() {
	http.HandleFunc("/", Handle.HomeHandle)
	http.HandleFunc("/Artist", Handle.ArtistsHandle)
	http.HandleFunc("/Search", Handle.SearchHandle)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
