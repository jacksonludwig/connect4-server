package game

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jacksonludwig/connect4-server/game/handlers"
)

type ClientActionRequest struct {
	Category string      `json:"category"`
	Data     interface{} `json:"data"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkRequestOrigin,
}

// TODO disallow wrong origin
func checkRequestOrigin(req *http.Request) bool {
	return true
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

		var actionReq ClientActionRequest
		json.Unmarshal([]byte(msg), &actionReq)

		switch actionReq.Category {
		case "CreateGame":
			handlers.CreateGame(actionReq.Data)
			break

		default:
			break
		}

		fmt.Printf("message received: %s\n", msg)
		fmt.Printf("category: %s", actionReq.Category)
	}
}
