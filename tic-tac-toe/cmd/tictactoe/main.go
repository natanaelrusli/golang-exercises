package main

import (
	"os"

	"github.com/natanaelrusli/tic-tac-toe/internal/tictactoe"
)

func main() {
	game := tictactoe.NewTicTacToe()
	userInput := tictactoe.GetUserInput(os.Stdin)

	if userInput != "" {
		game.StartGame(userInput)
	}
}
