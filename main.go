package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/jacksonludwig/connect4-server/game"
)

const ADDR = "localhost:9001"

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/connect", game.Connect)

	fmt.Printf("listening on %s\n", ADDR)

	log.Fatal(http.ListenAndServe(ADDR, router))
}
