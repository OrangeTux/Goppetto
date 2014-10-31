package app

import (
	"code.google.com/p/go.net/websocket"
	"log"
)

func webHandler(ws *websocket.Conn) {
	var msg string

	client := ws.Request().RemoteAddr
	log.Println("Client connected: ", client)

	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("Websocket Disconnected waiting", err.Error())
			return
		}

		log.Printf("%v -> %v", client, msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			log.Println("Could not send message to ", client, err.Error())
		}
		log.Printf("%v <- %v", client, msg)
	}
}
