package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// TODO disallow wrong origin
func checkRequestOrigin(req *http.Request) bool {
  return true
}

func connect(writer http.ResponseWriter, req *http.Request) {
  upgrader.CheckOrigin = checkRequestOrigin

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

func main() {
  router := mux.NewRouter()

  router.HandleFunc("/connect", connect)

  log.Fatal(http.ListenAndServe(":9100", router))
}
