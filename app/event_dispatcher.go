package app

import (
	"encoding/json"
)

// Jsonified EventMessage's are send over sockets. EventMessage contains the
// name of the event with 'event' field. EventMessage contains data of
// event in a map in field 'data'.
type EventMessage struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}

type EventDispatcher struct {
	callbacks map[string][]func(*EventMessage) *EventMessage
}

// Bind a callback. Multiple callbacks can be bound to 1 event. A callback
// takes en EventMessages as parameter and returns an EventMessage.
func (ed EventDispatcher) Bind(e string, cb func(*EventMessage) *EventMessage) {
	if _, present := ed.callbacks[e]; !present {
		ed.callbacks[e] = make([]func(*EventMessage) *EventMessage, 0)
	}
	ed.callbacks[e] = append(ed.callbacks[e], cb)
}

// Fire callbacks when event arrives. Callbacks are fired inside go routine and
// are therefore asynchronous.
func (e EventDispatcher) Dispatch(emsg *EventMessage) {
	for _, callback := range e.callbacks[emsg.Event] {
		go callback(emsg)
	}
}

// Listen on channel for incomming messages and fires bound callbacks
// when message arrives.
func (e EventDispatcher) Listen(messags chan string) {
	for {
		msg := <-messags
		if len(msg) > 0 {
			em := &EventMessage{}

			json.Unmarshal([]byte(msg), em)
			e.Dispatch(em)
		}
	}
}
