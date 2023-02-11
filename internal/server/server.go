package server

import (
	"embed"
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed index.html
var index string

//go:embed css/style.css
var content embed.FS

var port = 33400

func Start() {
	log.Printf("INFO: starting enigma GUI on http://localhost:%d", port)
	http.HandleFunc("/", handleIndex)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(index))
	if err != nil {
		log.Println(err)
	}
}
