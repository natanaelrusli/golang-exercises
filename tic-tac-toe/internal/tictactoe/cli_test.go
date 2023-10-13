package tictactoe

import (
	"strings"
	"testing"
)

// TestGetUserInput tests the GetUserInput function
func TestGetUserInput(t *testing.T) {
	testCases := []struct {
		input  string
		expect string
	}{
		{"123456789\n", "123456789"},
		{"XXX\n", ""},
		{"\n", ""},
	}

	for _, tc := range testCases {
		stdin := strings.NewReader(tc.input)
		result := GetUserInput(stdin)

		// Check the result
		if result != tc.expect {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expect, result)
		}
	}
}

func TestDisplayBoard(t *testing.T) {
	testCases := []struct {
		given string
		want  string
	}{
		{"XXXOOOXXX", "\nGAME BOARD\n" +
			"X | X | X\n" +
			"O | O | O\n" +
			"X | X | X\n"},
	}

	for _, tc := range testCases {
		game := NewTicTacToe()
		game.SplitIntoBoard(tc.given)
		result := game.DisplayBoard()

		// Check the result
		if tc.want != result {
			t.Errorf("Input: %s, Expected: %s, Got: %s", tc.given, tc.want, result)
		}
	}
}
