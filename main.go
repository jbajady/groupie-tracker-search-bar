package main

import (
	"fmt"
	"net/http"

	Handle "GroupieTracker/FuncHandle"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handle.HomeHandle)
	mux.HandleFunc("/Artist", Handle.ArtistsHandle)
	mux.HandleFunc("/Search", Handle.SearchHandle)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
