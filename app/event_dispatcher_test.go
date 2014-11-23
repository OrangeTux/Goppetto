package app

import (
	"testing"
)

var ed EventDispatcher

func init() {
	ed = EventDispatcher{
		callbacks: make(map[string][]func(*EventMessage) *EventMessage),
	}
}

func TestBind(t *testing.T) {
	someEvent := func(e *EventMessage) (emsg *EventMessage) {
		return e
	}

	someOtherEvent := func(e *EventMessage) (emsg *EventMessage) {
		return e
	}

	ed.Bind("some_event", someEvent)
	ed.Bind("some_event", someOtherEvent)

	if len(ed.callbacks["some_event"]) != 2 {
		t.Errorf("EventDispatcher should have 2 callbacks bound to 'some_event' got %v", len(ed.callbacks["some_event"]))
	}
}

func TestDispatch(t *testing.T) {
	msg := EventMessage{Event: "some_event"}
	chnl := make(chan bool)

	isCalled := false
	someEvent := func(e *EventMessage) (emsg *EventMessage) {
		isCalled = true
		chnl <- true

		return e
	}

	ed.Bind("some_event", someEvent)
	ed.Dispatch(&msg)

	// Wait for the someEvent to be called.
	<-chnl

	if isCalled == false {
		t.Error("Method has not been dispatched.")
	}
}

func TestListen(t *testing.T) {
	msg := `{"event": "some_event", "data": {"pin_id": 1, "state": 0}}`
	isExecuted := false
	done := make(chan bool)
	messages := make(chan string, 1)

	someEvent := func(e *EventMessage) (emsg *EventMessage) {
		isExecuted = true
		done <- true
		return e
	}

	ed.Bind("some_event", someEvent)
	go ed.Listen(messages)
	messages <- msg

	<-done

	if isExecuted == false {
		t.Error("someEvent has not been dispatched.")
	}

	close(messages)
}

// Global EventDispatcher instance
// TestListen
