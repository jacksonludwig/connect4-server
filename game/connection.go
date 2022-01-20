package game

import (
	"sync"
	"github.com/gorilla/websocket"
)

type Connection struct {
	Socket *websocket.Conn
	mutex  sync.Mutex
}

func (c *Connection) Send(message string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.Socket.WriteJSON(message)
}
