package goppetto

import (
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketManager struct {
	Connections map[net.Addr]*websocket.Conn
}

func (wsm WebSocketManager) ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Could not handle new websocket connection.")
	}

	addr := conn.RemoteAddr()
	wsm.Connections[addr] = conn

	log.Printf("%s connected.", addr)
}
