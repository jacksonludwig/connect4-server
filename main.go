package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"
)

func main() {
  router := mux.NewRouter()

  // TODO
  router.HandleFunc("/createGame", createGame)

  log.Fatal(http.ListenAndServe(":9100", router))
}
