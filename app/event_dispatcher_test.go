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

func TestBind(t *testing.T) {
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

func TestDispatch(t *testing.T) {
	msg := EventMessage{event: "some_event"}
	chnl := make(chan bool)

	isCalled := false
	someEvent := func(e EventMessage) (emsg EventMessage) {
		isCalled = true
		chnl <- true

		return e
	}

	ed.Bind("some_event", someEvent)
	ed.dispatch(msg)

	// Wait for the callback to be called.
	<-chnl

	if isCalled == false {
		t.Error("Method has not been dispatched.")
	}
}

// Global EventDispatcher instance
// TestListen
