package tictactoe

import (
	"fmt"
	"io"
	"strings"
)

func GetUserInput(reader io.Reader) string {
	var userInput = ""
	var attempts = 0
	var maxAttempts = 3

	for userInput == "" {

		fmt.Println("Sample Input: ")

		fmt.Fscanln(reader, &userInput)
		strArr := strings.Split(userInput, "")

		if len(strArr) != 9 {
			userInput = ""
			attempts += 1

			if attempts == maxAttempts {
				fmt.Println("Exceeded maximum attempts. Exiting.")
				break
			}

			fmt.Println("please input exactly 9 characters")
		}
	}

	return userInput
}

func (game *TicTacToe) DisplayBoard() string {
	resultString := "\nGAME BOARD\n"

	for _, v := range game.board {
		resultString += v[0] + " | " + v[1] + " | " + v[2] + "\n"
	}

	return resultString
}
