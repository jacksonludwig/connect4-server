module github.com/jacksonludwig/connect4-server/game

go 1.16

require github.com/gorilla/websocket v1.4.2
replace github.com/jacksonludwig/connect4-server/game/handlers => ./game/handlers
