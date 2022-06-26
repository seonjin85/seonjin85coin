package p2p

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/seonjin85/seonjin85coin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// port :3000 will upgared the request from :4000
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// http conn --> websocket conn
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)

	openPort := r.URL.Query().Get("openPort")
	result := strings.Split(r.RemoteAddr, ":")
	initPeer(conn, result[0], openPort)
}

func AddPeer(address, port, openPort string) {
	// from :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
