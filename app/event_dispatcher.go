package app

// Jsonified EventMessage's are send over sockets. EventMessage contains the
// name of the event with 'event' field. EventMessage contains data of
// event in a map in field 'data'.
type EventMessage struct {
	event string
	data  map[string]interface{}
}

type EventDispatcher struct {
	messages  chan string
	callbacks map[string][]func(EventMessage) EventMessage
}

// Bind a callback. Multiple callbacks can be bound to 1 event. A callback
// takes en EventMessages as parameter and returns an EventMessage.
func (ed EventDispatcher) Bind(e string, cb func(EventMessage) EventMessage) {
	if _, present := ed.callbacks[e]; !present {
		ed.callbacks[e] = make([]func(EventMessage) EventMessage, 0)
	}
	ed.callbacks[e] = append(ed.callbacks[e], cb)
}

func (e EventDispatcher) dispatch(msg EventMessage) {
	for _, callback := range e.callbacks[msg.event] {
		callback(msg)
	}
}
