package app

import (
	"testing"
)

var ed EventDispatcher
var isCalled bool
var done = make(chan bool)

var f = func(emsg *EventMessage) *EventMessage {
	isCalled = true
	done <- true

	return emsg
}

func setUp() {
	ed = EventDispatcher{
		callbacks: make(map[string][]func(*EventMessage) *EventMessage),
	}
	isCalled = false
}

func TestBind(t *testing.T) {
	setUp()
	ed.Bind("some_event", f)
	ed.Bind("some_event", f)

	if len(ed.callbacks["some_event"]) != 2 {
		t.Errorf("EventDispatcher should have 2 callbacks bound to 'some_event' got %v", len(ed.callbacks["some_event"]))
	}
}

func TestListen(t *testing.T) {
	setUp()
	msg := `{"event": "some_event", "data": {"pin_id": 1, "state": 0}}`
	messages := make(chan string, 1)

	ed.Bind("some_event", f)
	go ed.Listen(messages)
	messages <- msg

	<-done

	if isCalled == false {
		t.Error("someEvent has not been dispatched.")
	}
}

func TestDispatch(t *testing.T) {
	setUp()
	msg := EventMessage{Event: "some_event"}

	ed.Bind("some_event", f)
	ed.Dispatch(&msg)

	// Wait for the someEvent to be called.
	<-done

	if isCalled == false {
		t.Error("Method has not been dispatched.")
	}
}
