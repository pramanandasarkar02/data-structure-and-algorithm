package main

import "C"
import "fmt"

var (
	N     int
	board [][]int
)

func initBoard() {
	board = make([][]int, N)
	for i := 0; i < N; i++ {
		board[i] = make([]int, N)
	}
}

func isSafe(row int, col int) bool {
	// Check row on left side
	for j := 0; j < col; j++ {
		if board[row][j] == 1 {
			return false
		}
	}

	// Check upper diagonal on left side
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	// Check lower diagonal on left side
	for i, j := row, col; i < N && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func solve(col int) bool {
	if col >= N {
		return true
	}

	for row := 0; row < N; row++ {
		if isSafe(row, col) {
			board[row][col] = 1
			if solve(col + 1) {
				return true
			}
			board[row][col] = 0 
		}
	}
	return false
}

func solveNQueens() [][]string {
	N = 8
	initBoard()
	solve(0)

	
	result := make([][]string, N)
	for i := 0; i < N; i++ {
		result[i] = make([]string, N)
		for j := 0; j < N; j++ {
			if board[i][j] == 1 {
				result[i][j] = "Q "
			} else {
				result[i][j] = ". "
			}
		}
	}
	return result
}

func main() {
	solution := solveNQueens()
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Print(solution[i][j] + " ")
		}
		fmt.Println()
	}
}