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

	done := make(chan bool)

	wsm := WebSocketManager{
		make(map[string]*websocket.Conn),
		make([]func(*websocket.Conn) *websocket.Conn, 0),
		make([]func([]byte, *websocket.Conn) ([]byte, *websocket.Conn), 0),
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

		Convey("When an `OnMessage` callback has been bound", func() {
			wsm.OnMessage(func(msg []byte, conn *websocket.Conn) ([]byte, *websocket.Conn) {
				done <- true
				return msg, conn
			})
			Convey("And a client sends a message", func() {
				conn.WriteMessage(websocket.TextMessage, []byte{'H', 'e', 'l', 'l', 'o'})
				Convey("Then the callback must be called", func() {
					<-done
					// When test reaches to this point is has callback has been called test passed.
					So(1, ShouldEqual, 1)
				})
			})
		})

		Convey("When an `OnDisconnect` callback has been bound", func() {
			wsm.OnDisconnect(func(conn *websocket.Conn) *websocket.Conn {
				done <- true
				return conn
			})

			Convey("And a connected client disconnects", func() {
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
	})
}
