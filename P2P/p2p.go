package p2p

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/seonjin85/seonjin85coin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// http conn --> websocket conn
	_, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
}
