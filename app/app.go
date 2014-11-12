package app

import (
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/specs", specs)

	http.Handle("/api", websocket.Handler(SocketHandler))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func Run() {
	log.Printf("Start Goppetto on http://0.0.0.0:9999.")
	log.Fatal(http.ListenAndServe(":9999", nil))
}
