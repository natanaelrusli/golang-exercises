package tictactoe

type TicTacToe struct {
	board        [3][3]string
	possibleCond [8][3]string
}

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		board: [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}},
	}
}
