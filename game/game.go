package game

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var sockets = make(map[string]*Connection)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Connect(writer http.ResponseWriter, req *http.Request) {
	connection, err := upgrader.Upgrade(writer, req, nil)

	if err != nil {
		fmt.Printf("error upgrading connection: %s\n", err.Error())
		return
	}

	defer connection.Close()

	// main event loop
	for {
		_, msg, err := connection.ReadMessage()

		if err != nil {
			fmt.Printf("error reading message: %s\n", err.Error())
			break
		}

		fmt.Printf("message received: %s\n", msg)
	}
}
