package wshub

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type WSHub struct {
	clients []*websocket.Conn
	sync.Mutex
}

func (hub *WSHub) Add(c *websocket.Conn) {
	hub.Lock()
	hub.clients = append(hub.clients, c)
	hub.Unlock()
}

func (hub *WSHub) Send(msg interface{}) {
	msgData, _ := json.Marshal(msg)

	hub.Lock()
	for _, c := range hub.clients {
		_ = c.WriteMessage(websocket.TextMessage, msgData)
	}
	hub.Unlock()
}

func (hub *WSHub) Close() error {
	hub.Lock()
	for _, c := range hub.clients {
		_ = c.Close()
	}
	hub.Unlock()
	return nil
}
