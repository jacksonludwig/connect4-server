module github.com/jacksonludwig/connect4-server

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	github.com/jacksonludwig/connect4-server/game v0.0.0-20220118195722-558f2c1c63a9
)

replace github.com/jacksonludwig/connect4-server/game => ./game
