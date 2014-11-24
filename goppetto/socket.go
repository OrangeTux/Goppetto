package goppetto

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

// A channel where event messages are passed through.
var messages = make(chan string)

// A map with all connections of all connected clients stored by string of
// ip and socket.
var connections = make(map[string]*websocket.Conn)

func init() {
	go keepTalking(messages)
}

// Keep writing messages from messages channel to sockets of all connection
// clients.
func keepTalking(messages chan string) {
	for {
		msg := <-messages
		log.Printf("Amount of connected clients: %d.", len(connections))

		for addr, client := range connections {
			err := talk(client, msg)
			if err != nil {
				log.Printf("Could not send %s to %s: %s\n", msg, addr, err.Error())
			} else {
				log.Printf("Send %v to %s.\n", msg, addr)
			}
		}
	}
}

// Write message on socket.
func talk(con *websocket.Conn, msg string) (err error) {
	return websocket.Message.Send(con, msg)
}

// Keeps listening on socket for incomming event messages and send all incoming
// event messages to messages channel.
func keepListening(con *websocket.Conn, messages chan string) {
	for {
		msg, err := listen(con)
		addr := con.Request().RemoteAddr
		if err != nil {
			log.Printf("Unable to read message from %s: %s\n", addr, err.Error())
			return
		} else {
			log.Printf("Read message %s from %s.", msg, addr)
			messages <- msg
		}
	}
}

// Listen on socket for incoming event messages. Return message on receive.
func listen(con *websocket.Conn) (msg string, err error) {
	err = websocket.Message.Receive(con, &msg)

	return msg, err
}

// Append connection to list of all connections and start listening on that
// connection for messages.
func SocketHandler(con *websocket.Conn) {
	addr := con.Request().RemoteAddr
	log.Printf("%s connected.", addr)

	// Append connection to map with connections.
	connections[addr] = con

	// Listen for incoming messages.
	keepListening(con, messages)

	log.Printf("%s disconnected.", addr)
	delete(connections, addr)
}
