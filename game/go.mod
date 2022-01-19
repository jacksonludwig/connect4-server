module github.com/jacksonludwig/connect4-server/game

go 1.16

require (
	github.com/gorilla/websocket v1.4.2
	github.com/jacksonludwig/connect4-server/game/handlers v0.0.0-20220119172957-4b50ca2d0002
)

replace github.com/jacksonludwig/connect4-server/game/handlers => ./handlers
