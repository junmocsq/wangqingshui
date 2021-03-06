package controllers

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"wangqingshui/services/ws"
)

type WebSocketController struct {
	baseController
}

func (w *WebSocketController) NestPrepare() {

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 120 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocket 加入
// @router / [get]
func (w *WebSocketController) websocket() {
	wsConn, err := upgrader.Upgrade(w.Ctx.ResponseWriter, w.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	ws.NewWs().Conn(wsConn, w.ip)
}
