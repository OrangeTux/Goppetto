package goppetto

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEventDispatcher(t *testing.T) {
	Convey("Given I have a EventDispatcher", t, func() {
		i := 0
		done := make(chan bool)

		em := EventMessage{"pin_state", make(map[string]interface{})}

		ed := EventDispatcher{make(map[string][]func(*EventMessage) *EventMessage)}

		Convey("When I bind a callback to an event", nil)

		Convey("And that the EventDispatcher receives this event", func() {
			ed.Bind("pin_state", func(e *EventMessage) *EventMessage {
				i += 1
				done <- true

				return e
			})

			ed.Dispatch(&em)

			Convey("Then the callback must be executed.", func() {
				<-done
				So(i, ShouldEqual, 1)
			})
		})

		Convey("When I bind multiple callbacks to an event", nil)

		Convey("And the EventDispatcher receives this event", func() {
			signal := make(chan bool)

			ed.Bind("pin_state", func(e *EventMessage) *EventMessage {
				i += 1
				signal <- true
				return e
			})

			// This callback receives a value in channel `signal`. This value
			// can only be send by the previous callback. If both callbacks
			// where executed seqentually this construct would cause a dead
			// lock.
			ed.Bind("pin_state", func(e *EventMessage) *EventMessage {
				i += 1
				<-signal
				done <- true
				return e
			})

			ed.Dispatch(&em)

			Convey("Then the callbacks must be executed in parallel.", func() {
				<-done
				So(i, ShouldEqual, 2)
			})
		})
	})
}
