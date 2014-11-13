package app

import (
	"testing"
)

func TestEventsShouldBeDispatched(t *testing.T) {
	ed := EventDispatcher{make(map[string]func(EventMessage) EventMessage)}
	msg := EventMessage{event: "some_event"}

	isCalled := false

	x := func(e EventMessage) (emsg EventMessage) {
		isCalled = true

		return e
	}

	ed.callbacks["some_event"] = x

	ed.dispatch(msg)

	if isCalled == false {
		t.Error("Fails")
	}
}
