package goppetto

import (
	"encoding/json"
)

// An EventMessage is a simple struct type which containts 2 attributes:
// `Event` and `Data`. `Event` contains a string with the name of the event.
// `Data` is a map with string keys, the values can be anything.
type EventMessage struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}

// The EventDispatcher can be used to dispatch events on incomming messages.
//
// Create an EventDispatcher.
//
//	var ed = EventDispatcher{
// 		callbacks: make(map[string][]func(*EventMessage) *EventMessage),
//	}
//
// Now create an callback which listens for an event with the name
// `set_direction`. A callback receives a pointer to EventMessage as input and
// must return this pointer. The pointer must be returned so we can chain
// callbacks.
//
// In our example our callback receives a struct which data map contains a key
// `direction` which holds an string with like `north` or `south`.
//
//	func setDirection(emsg *EventMessage) *EventMessage {
//		log.Println(emsg.Data['direction'])
//		return emsg
//	}
//
// Now bind this callback to the `set_direction`. You can bind multiple
// callbacks to 1 event. They will be executed asynchronously.
//
//	ed.Bind('set_direction', setDirection)
//
// The only thing you've left to do is let the EventDispatcher listen on a
// channel.
//
// 	messages := make(chan string)
//	ed.Listen(messages)
//
// You can send messages through the channel. Messages are strings following
// the JSON syntax with the format:
//
//	{
//		"event": <event_name>,
//		"data": {
//			<first_attribute>: <some_value>,
//			<second_attribute>: <some_other_value>,
//			...
//		}
//	}
//
// 	messages <- `{"event": "set_direction", "data": {"direction": "north"}}`
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
func (e EventDispatcher) Listen(messages chan []byte) {
	for msg := range messages {
		if len(msg) > 0 {
			em := &EventMessage{}

			json.Unmarshal(msg, em)
			e.Dispatch(em)
		}
	}
}
