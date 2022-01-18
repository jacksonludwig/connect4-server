package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"
)

const ADDR = "localhost:9001"

var upgrader = websocket.Upgrader{
  CheckOrigin: checkRequestOrigin,
}

// TODO disallow wrong origin
func checkRequestOrigin(req *http.Request) bool {
  return true
}

func connect(writer http.ResponseWriter, req *http.Request) {
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

  fmt.Printf("listening on %s\n", ADDR)

  log.Fatal(http.ListenAndServe(ADDR, router))
}
