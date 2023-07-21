package main

import (
	"fmt"
	"strings"
)

// TicTacToe struct represents the game board and the current player.
type TicTacToe struct {
	board        [3][3]string //board structure
	currentPlayer string //player X o
}

// NewTicTacToe creates a new instance of TicTacToe game.
func NewTicTacToe() *TicTacToe {
	game := TicTacToe{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			game.board[i][j] = " "
		}
	}
	game.currentPlayer = "X"
	return &game
}

// PrintBoard prints the current state of the game board.
func (game *TicTacToe) PrintBoard() {
	fmt.Println("---------")
	for i := 0; i < 3; i++ {
		fmt.Printf("| %s | %s | %s |\n", game.board[i][0], game.board[i][1], game.board[i][2])
		fmt.Println("---------")
	}
}

// MakeMove makes a move at the specified row and column for the current player.
func (game *TicTacToe) MakeMove(row, col int) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("Invalid move. Row and column should be between 0 and 2")
	}

	if game.board[row][col] != " " {
		return fmt.Errorf("Invalid move. Cell already occupied")
	}

	game.board[row][col] = game.currentPlayer
	game.currentPlayer = switchPlayer(game.currentPlayer)
	return nil
}

// switchPlayer switches the current player from X to O, or vice versa.
func switchPlayer(player string) string {
	if player == "X" {
		return "O"
	}
	return "X"
}

// CheckWin checks if the current player has won the game.
func (game *TicTacToe) CheckWin() bool {
	for i := 0; i < 3; i++ {
		// Check rows
		if strings.Join(game.board[i][:], "") == "XXX" || strings.Join(game.board[i][:], "") == "OOO" {
			return true
		}

		// Check columns
		if game.board[0][i] == game.currentPlayer && game.board[1][i] == game.currentPlayer && game.board[2][i] == game.currentPlayer {
			return true
		}
	}

	// Check diagonals
	if (game.board[0][0] == game.currentPlayer && game.board[1][1] == game.currentPlayer && game.board[2][2] == game.currentPlayer) ||
		(game.board[0][2] == game.currentPlayer && game.board[1][1] == game.currentPlayer && game.board[2][0] == game.currentPlayer) {
		return true
	}

	return false
}
// entry point 
func main() {
	game := NewTicTacToe()
	fmt.Println("Welcome to Tic Tac Toe!")

	for {
		game.PrintBoard()
		fmt.Printf("Player %s's turn. Enter row and column (0-2) separated by space: ", game.currentPlayer)

		var row, col int
		_, err := fmt.Scanf("%d %d", &row, &col)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		err = game.MakeMove(row, col)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if game.CheckWin() {
			fmt.Printf("Player %s wins!\n", game.currentPlayer)
			break
		}

		// Check for a draw
		draw := true
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if game.board[i][j] == " " {
					draw = false
					break
				}
			}
		}
		if draw {
			fmt.Println("It's a draw!")
			break
		}
	}

	game.PrintBoard()
}
