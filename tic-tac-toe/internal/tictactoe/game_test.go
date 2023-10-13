package tictactoe

import (
	"fmt"
	"testing"
)

func TestSplitIntoBoard(t *testing.T) {
	given := "XOXOOOXOO"
	want := [3][3]string{
		{"X", "O", "X"},
		{"O", "O", "O"},
		{"X", "O", "O"},
	}

	game := NewTicTacToe()
	game.SplitIntoBoard(given)

	res := game.board
	if res != want {
		t.Fatalf("SplitIntoBoard() = %q want %#q", res, want)
	}
}

func TestMoveToPossibleCond(t *testing.T) {
	given := "XOXOOOXOO"
	want := [8][3]string{
		{"X", "O", "O"},
		{"X", "O", "X"},
		{"X", "O", "X"},
		{"O", "O", "O"},
		{"X", "O", "O"},
		{"X", "O", "X"},
		{"O", "O", "O"},
		{"X", "O", "O"},
	}

	game := NewTicTacToe()
	game.SplitIntoBoard(given)
	game.MapToPossibleCond()

	res := game.possibleCond
	if res != want {
		t.Fatalf("MapToPossibleCond() = %q want %#q", res, want)
	}
}

func TestCheckGameBoard(t *testing.T) {
	testCases := []struct {
		given string
		want  string
	}{
		{
			given: "XOXOOOXOO",
			want:  "invalid game board",
		},
		{
			given: "XOXXOOXXO",
			want:  "X wins!",
		},
		{
			given: "XOOXOXOXO",
			want:  "O wins!",
		},
		{
			given: "OXOXOXXOX",
			want:  "Its a draw!",
		},
		{
			given: "XOXX--O--",
			want:  "Game still in progress!",
		},
	}

	game := NewTicTacToe()

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("TestCheckGameBoard given %q want %q", tc.given, tc.want), func(t *testing.T) {
			game.SplitIntoBoard(tc.given)
			game.MapToPossibleCond()

			res := game.CheckGameBoard()

			if res != tc.want {
				t.Fatalf("TestCheckGameBoard('XOXOOOXOO') = %q want %#q", res, tc.want)
			}
		})
	}
}

func TestStartGame(t *testing.T) {
	game := NewTicTacToe()

	testCases := []struct {
		given string
		want  string
	}{
		{
			given: "XOXOOOXOO",
			want:  "invalid game board",
		},
		{
			given: "XOXXOOXXO",
			want:  "X wins!",
		},
		{
			given: "XOOXOXOXO",
			want:  "O wins!",
		},
		{
			given: "OXOXOXXOX",
			want:  "Its a draw!",
		},
		{
			given: "XOXX--O--",
			want:  "Game still in progress!",
		},
	}

	for _, tc := range testCases {
		result := game.StartGame(tc.given)

		// Check the result
		if result != tc.want {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.given, tc.want, result)
		}
	}
}
