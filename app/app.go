package app

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

}

func Run() {
	log.Printf("Start Goppetto on http://127.0.0.1:9999.")
	log.Fatal(http.ListenAndServe("localhost:9999", nil))
}
