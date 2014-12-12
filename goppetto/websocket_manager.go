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

// The WebSocketManager keeps track of all connected clients and offers ways to get updated when a client sends a
// message of disconnects.
//
// The ConnectionHandler is handler for incoming WS requests.
//
//  wsm := WebSocketManager{
//      make(map[string]*websocket.Conn),
//      make([]func(*websocket.Conn) *websocket.Conn, 0),
//      make([]func([]byte, *websocket.Conn) ([]byte, *websocket.Conn), 0),
//  }
//
//  http.HandleFunc("/ws", wsm.ConnectionHandler)
//
// Use OnMessage to register a callback that is fired when a client sends a message. The callback must receive
// a slice of bytes and an pointer to websocket.Conn and must return both objects to make chaining of callbacks
// possible. Multiple callbacks can be bound and are executed in parallel.
//
//  messageCallback := func(msg []byte, conn *websocket.Conn) ([]byte, *websocket.Conn) {
//      fmt.Println("Client %s send %s", conn.RemoteAddr().String(), msg)
//  }
//
//  wsm.OnMessage(messageCallback)
//
// Use OnDisconnect to register a callback that is fired when a client disconnects. The callback must receive a
// pointer to a websocket.Conn and must return this pointer to. Multiple callbacks are executed in parallel.
//
//  disconnectCallback := func(conn *websocket.Conn) *websocket.Conn {
//      fmt.Println("Client %s disconnected.", conn.RemoteAddr().String)
//  }
//
//  wsm.OnDisconnect(disconnectCallback)
type WebSocketManager struct {
	// A map with all connected clients.
	Connections map[string]*websocket.Conn
	// A slice with all callbacks that are executed when a client disconnects.
	disconnectCallbacks []func(*websocket.Conn) *websocket.Conn
	// A slice will all callbacks that are executed when a client sends a message.
	messageCallbacks []func([]byte, *websocket.Conn) ([]byte, *websocket.Conn)
}

// A handler that can be used create new conections.
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

// Register an callback that will be executed when a clients disconnects.
func (wsm *WebSocketManager) OnDisconnect(f func(*websocket.Conn) *websocket.Conn) {
	wsm.disconnectCallbacks = append(wsm.disconnectCallbacks, f)
}

// Register an callback that will be executed when a clients sends a message.
func (wsm *WebSocketManager) OnMessage(f func(msg []byte, conn *websocket.Conn) ([]byte, *websocket.Conn)) {
	wsm.messageCallbacks = append(wsm.messageCallbacks, f)
}

// Listen for message on connection. When reading from connection fails, `disconnect` is called. When no error appears
// `message` is called.
func (wsm *WebSocketManager) listen(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			wsm.disconnect(conn)
			return
		}

		wsm.message(msg, conn)
	}
}

// Close connection, remove the connection from Connections and execute `OnDisconnect`.
func (wsm WebSocketManager) disconnect(conn *websocket.Conn) {
	conn.Close()
	delete(wsm.Connections, conn.RemoteAddr().String())
	log.Printf("%s disconnected.", conn.RemoteAddr().String())
	for _, f := range wsm.disconnectCallbacks {
		go f(conn)
	}
}

// Execute `OnMessage` callbacks.
func (wsm WebSocketManager) message(msg []byte, conn *websocket.Conn) ([]byte, *websocket.Conn) {
	log.Printf("Received `%s` from %s.", msg, conn.RemoteAddr())
	for _, f := range wsm.messageCallbacks {
		go f(msg, conn)
	}

	return msg, conn
}
