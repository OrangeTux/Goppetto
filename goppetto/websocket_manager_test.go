package goppetto

import (
	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSocketManager(t *testing.T) {

	wsm := WebSocketManager{make(map[net.Addr]*websocket.Conn)}
	ts := httptest.NewServer(http.HandlerFunc(wsm.ConnectionHandler))
	defer ts.Close()

	url, _ := url.Parse(ts.URL)
	url.Scheme = "ws"

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	defer conn.Close()

	log.Print(conn.LocalAddr())

	if err != nil {
		log.Print(ts.URL)
		log.Fatal(err)
	}

	Convey("Given a WebSocketManager", t, func() {

		Convey("When a client connects", func() {

			Convey("Then the connection must be saved.", func() {
				So(wsm.Connections[conn.LocalAddr()], ShouldHaveSameTypeAs, conn)
			})
		})

		Convey("When I bind a callback on incomming messages", nil)
		Convey("And a client sends a message", func() {
			Convey("Then the callback must be called", nil)
		})

		Convey("When a client sends a message", func() {

			Convey("Then message must be processed.", nil)

		})

		Convey("When a client disconnects", func() {

			Convey("Then the connections must be flushed.", nil)

		})
	})
}
