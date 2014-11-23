package app

import (
	"testing"
)

var ed EventDispatcher

func init() {
	ed = EventDispatcher{
		messages:  make(chan string),
		callbacks: make(map[string][]func(EventMessage) EventMessage),
	}
}

//func TestEventsShouldBeDispatched(t *testing.T) {
//ed := EventDispatcher{make(map[string]func(EventMessage) EventMessage)}
//msg := EventMessage{event: "some_event"}

//isCalled := false

//x := func(e EventMessage) (emsg EventMessage) {
//isCalled = true

//return e
//}

//ed.callbacks["some_event"] = x

//ed.dispatch(msg)

//if isCalled == false {
//t.Error("Fails")
//}
//}

func TestBindCallback(t *testing.T) {
	someEvent := func(e EventMessage) (emsg EventMessage) {
		return e
	}

	someOtherEvent := func(e EventMessage) (emsg EventMessage) {
		return e
	}

	ed.Bind("some_event", someEvent)
	ed.Bind("some_event", someOtherEvent)

	if len(ed.callbacks["some_event"]) != 2 {
		t.Errorf("EventDispatcher should have 2 callbacks bound to 'some_event' got %v", len(ed.callbacks["some_event"]))
	}
}

// Global EventDispatcher instance
// TestDispatch
// TestListen
