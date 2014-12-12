package goppetto

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebSocketManager struct {
	Connections         map[string]*websocket.Conn
	disconnectCallbacks []func(*websocket.Conn) *websocket.Conn
}

func (wsm *WebSocketManager) ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Could not handle new websocket connection.")
	}

	addr := conn.RemoteAddr().String()
	wsm.Connections[addr] = conn

	log.Printf("%s connected.", addr)

	wsm.listen(conn)
}

// Bind a callback when client disconnects.
func (wsm *WebSocketManager) OnDisconnect(f func(*websocket.Conn) *websocket.Conn) {
	wsm.disconnectCallbacks = append(wsm.disconnectCallbacks, f)
}

// Listen for message on connection.
func (wsm *WebSocketManager) listen(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			wsm.disconnect(conn)
			return
		}

		log.Printf("Received %s from %s.", msg, conn.RemoteAddr())
	}
}

// Close connection, remove the connection from map and execute `OnDisconnect`
// callbacks.
func (wsm WebSocketManager) disconnect(conn *websocket.Conn) {
	conn.Close()
	delete(wsm.Connections, conn.RemoteAddr().String())
	for _, f := range wsm.disconnectCallbacks {
		go f(conn)
	}
}
