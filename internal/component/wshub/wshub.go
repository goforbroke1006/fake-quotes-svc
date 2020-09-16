package wshub

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type WSHub struct {
	clients []*websocket.Conn
	sync.RWMutex
}

func (hub *WSHub) Add(c *websocket.Conn) {
	hub.Lock()
	hub.clients = append(hub.clients, c)
	hub.Unlock()
}

func (hub *WSHub) Send(msg interface{}) {
	msgData, _ := json.Marshal(msg)

	hub.RLock()
	for _, c := range hub.clients {
		_ = c.WriteMessage(websocket.TextMessage, msgData)
	}
	hub.RUnlock()
}

func (hub *WSHub) Close() error {
	hub.Lock()
	for _, c := range hub.clients {
		_ = c.Close()
	}
	hub.Unlock()
	return nil
}
