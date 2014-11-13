package app

// Jsonified EventMessage's are send over sockets. EventMessage contains the
// name of the event with 'event' field. EventMessage contains data of
// event in a map in field 'data'.
type EventMessage struct {
	event string
	data  map[string]interface{}
}

type EventDispatcher struct {
	callbacks map[string]func(EventMessage) EventMessage
}

func (e EventDispatcher) dispatch(msg EventMessage) {
	for _, callback := range e.callbacks {
		callback(msg)
	}
}
