package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	tetrominoes, err := ParseTetrominoes(string(content))
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	board := solve(tetrominoes)
	printBoard(board)
}
