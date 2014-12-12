package goppetto

import (
	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSocketManager(t *testing.T) {

	wsm := WebSocketManager{
		make(map[string]*websocket.Conn),
		make([]func(*websocket.Conn) *websocket.Conn, 0),
	}

	ts := httptest.NewServer(http.HandlerFunc(wsm.ConnectionHandler))
	defer ts.Close()

	url, _ := url.Parse(ts.URL)
	url.Scheme = "ws"

	conn, _, _ := websocket.DefaultDialer.Dial(url.String(), nil)
	defer conn.Close()

	Convey("Given a WebSocketManager", t, func() {

		Convey("When a client connects", func() {

			Convey("Then the connection must be saved.", func() {
				So(wsm.Connections[conn.LocalAddr().String()], ShouldHaveSameTypeAs, conn)
			})
		})

		Convey("When I bind a callback on incomming messages", nil)
		Convey("And a client sends a message", func() {
			Convey("Then the callback must be called", nil)
		})

		Convey("When `OnDisconnect` callback has been bound", nil)
		Convey("And a connected client disconnects", func() {
			done := make(chan bool)
			f := func(conn *websocket.Conn) *websocket.Conn {
				done <- true
				return conn
			}
			wsm.OnDisconnect(f)

			conn.Close()

			Convey("Then the `OnDisconnect` callbacks must be fired.", func() {
				Convey("And the connection must be removed.", func() {

					<-done

					keys := make([]string, len(wsm.Connections))
					for key := range wsm.Connections {
						keys = append(keys, key)
					}
					So(conn.LocalAddr().String(), ShouldNotBeIn, keys)
				})
			})
		})
	})
}
