package main

import (
	"errors"
	"strings"
)

func ParseTetrominoes(content string) ([]Tetromino, error) {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	var tetrominoes []Tetromino

	blocks := strings.Split(content, "\n\n")

	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		if len(lines) != 4 {
			return nil, errors.New("invalid block height")
		}

		hashCount := 0

		for _, line := range lines {
			if len(line) != 4 {
				return nil, errors.New("invalid line width")
			}

			for _, char := range line {
				if char != '.' && char != '#' {
					return nil, errors.New("invalid character")
				}

				if char == '#' {
					hashCount++
				}
			}
		}

		if hashCount != 4 {
			return nil, errors.New("invalid hash count")
		}

		if !isValidShape(lines) {
			return nil, errors.New("invalid shape")
		}

		tetromino := createTetromino(lines, rune('A'+len(tetrominoes)))
		tetrominoes = append(tetrominoes, tetromino)
	}
	return tetrominoes, nil
}

func isValidShape(lines []string) bool {
	connections := 0

	for y, line := range lines {
		for x, char := range line {

			if char != '#' {
				continue
			}

			//Up
			if y > 0 && lines[y-1][x] == '#' {
				connections++
			}
			//Down
			if y < 3 && lines[y+1][x] == '#' {
				connections++
			}
			//Right
			if x < 3 && lines[y][x+1] == '#' {
				connections++
			}
			//Left
			if x > 0 && lines[y][x-1] == '#' {
				connections++
			}
		}
	}
	return connections == 6 || connections == 8
}

func createTetromino(lines []string, letter rune) Tetromino {
	var cells []Point

	minX := 4
	minY := 4

	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				cells = append(cells, Point{X: x, Y: y})

				if x < minX {
					minX = x
				}
				if y < minY {
					minY = y
				}
			}
		}
	}

	for i := range cells {
		cells[i].X -= minX
		cells[i].Y -= minY
	}

	return Tetromino{
		Cells:  cells,
		Letter: letter,
	}
}
