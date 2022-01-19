package handlers

import (
	"errors"
	"fmt"
)

type CreateGameData struct {
	Name string `json:"name"`
}

func CreateGame(data interface{}) error {
	createGameData, ok := data.(CreateGameData)

	if !ok {
		return errors.New(fmt.Sprintf("argument is invalid as create game data: %+v\n", data))
	}

	return nil
}
