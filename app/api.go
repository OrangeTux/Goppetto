package app

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

// A channel where messages are passed through.
var messages = make(chan string)

// A slice with all current connections.
var connections = make([]*websocket.Conn, 0)

func SocketHandler(ws *websocket.Conn) {
	var msg string

	client := ws.Request().RemoteAddr
	log.Println("Client connected: ", client)

	// Append connection to slice with connections.
	connections = append(connections, ws)
	//log.Printf("Amount of connections ", (len(connections)))
	log.Println(connections)

	go func() {
		for {
			select {
			case x := <-messages:
				{
					log.Println(x)
					for _, connection := range connections {
						if err := websocket.Message.Send(connection, x); err != nil {
							log.Println("Could not send message to ", client, err.Error())
						}
						log.Printf("%v <- %v", connection.Request().RemoteAddr, x)
					}
				}

			}
		}
	}()

	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("Websocket Disconnected waiting", err.Error())
			return
		}

		// Send the received message to the messages channel.
		log.Printf("%v -> %v", client, msg)
		messages <- msg
	}
}
