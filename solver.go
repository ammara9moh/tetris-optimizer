package main

import "fmt"

func solveBoard(board [][]rune, pieces []Tetromino, index int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board); x++ {
			if canPlace(board, piece, x, y) {
				placePiece(board, piece, x, y)

				if solveBoard(board, pieces, index+1) {
					return true
				}

				removePiece(board, piece, x, y)
			}
		}
	}

	return false
}

func solve(pieces []Tetromino) [][]rune {
	totalBlocks := len(pieces) * 4
	size := 1

	for size*size < totalBlocks {
		size++
	}

	for {
		board := createBoard(size)

		if solveBoard(board, pieces, 0) {
			return board
		}

		size++
	}
}

func createBoard(size int) [][]rune {
	board := make([][]rune, size)

	for y := 0; y < size; y++ {
		board[y] = make([]rune, size)
		for x := 0; x < size; x++ {
			board[y][x] = '.'
		}
	}

	return board
}

func canPlace(board [][]rune, piece Tetromino, startX, startY int) bool {
	size := len(board)

	for _, cell := range piece.Cells {
		x := startX + cell.X
		y := startY + cell.Y

		if x < 0 || x >= size || y < 0 || y >= size {
			return false
		}

		if board[y][x] != '.' {
			return false
		}
	}

	return true
}

func placePiece(board [][]rune, piece Tetromino, startX, startY int) {
	for _, cell := range piece.Cells {
		x := startX + cell.X
		y := startY + cell.Y

		board[y][x] = piece.Letter
	}
}

func removePiece(board [][]rune, piece Tetromino, startX, startY int) {
	for _, cell := range piece.Cells {
		x := startX + cell.X
		y := startY + cell.Y

		board[y][x] = '.'
	}
}

func printBoard(board [][]rune) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}
