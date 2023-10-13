package tictactoe

import (
	"fmt"
	"strings"
)

/*
possibleCondMapping := map[int]string{
	0: "Diagonal left top to right bottom",
	1: "Diagonal right bottom to left top",
	2: "Row 1",
	3: "Row 2",
	4: "Row 3",
	5: "Column 1",
	6: "Column 2",
	7: "Column 3",
}
*/

func (game *TicTacToe) StartGame(userInput string) string {
	game.SplitIntoBoard(userInput)

	board := game.DisplayBoard()
	fmt.Print(board)

	res := game.CheckGameBoard()

	fmt.Println(res)
	return res
}

func (game *TicTacToe) SplitIntoBoard(userInput string) [3][3]string {
	var userInputArr []string
	var boardMatrix [3][3]string

	userInputArr = strings.Split(userInput, "")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			boardMatrix[i][j] = userInputArr[i*3+j]
		}
	}

	game.board = boardMatrix
	game.MapToPossibleCond()

	return boardMatrix
}

func (game *TicTacToe) MapToPossibleCond() {
	possibleCond := [8][3]string{
		{game.board[0][0], game.board[1][1], game.board[2][2]},
		{game.board[0][2], game.board[1][1], game.board[2][0]},
		{game.board[0][0], game.board[0][1], game.board[0][2]},
		{game.board[1][0], game.board[1][1], game.board[1][2]},
		{game.board[2][0], game.board[2][1], game.board[2][2]},
		{game.board[0][0], game.board[1][0], game.board[2][0]},
		{game.board[0][1], game.board[1][1], game.board[2][1]},
		{game.board[0][2], game.board[1][2], game.board[2][2]},
	}

	game.possibleCond = possibleCond
}

func (game *TicTacToe) CheckGameBoard() string {
	inProgress := game.checkGameInProgress()

	if inProgress {
		return "Game still in progress!"
	}

	invalid := game.checkInvalidBoard()

	if invalid {
		return "invalid game board"
	}

	for _, val := range game.possibleCond {
		if val[0] == val[1] &&
			val[0] == val[2] {
			return fmt.Sprintf("%v wins!", val[0])
		}
	}

	return "Its a draw!"
}

func (game *TicTacToe) checkInvalidBoard() bool {
	xCount, oCount := 0, 0
	winCount := 0

	for _, val := range game.board {
		if val[0] == "X" {
			xCount++
		} else if val[0] == "O" {
			oCount++
		}

		if val[1] == "X" {
			xCount++
		} else if val[1] == "O" {
			oCount++
		}

		if val[2] == "X" {
			xCount++
		} else if val[2] == "O" {
			oCount++
		}
	}

	if xCount > 5 || oCount > 5 {
		return true
	}

	for _, val := range game.possibleCond {
		win := allElementsSame(val)

		if win {
			winCount += 1
		}
	}

	return winCount > 1
}

func (game *TicTacToe) checkGameInProgress() bool {
	for _, row := range game.board {
		for _, val := range row {
			if val != "O" && val != "X" {
				return true
			}
		}
	}

	return false
}

func allElementsSame(arr [3]string) bool {
	firstElement := arr[0]

	for _, element := range arr {
		if element != firstElement {
			return false
		}
	}

	return true
}
