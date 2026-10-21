package main

type Point struct {
	X int
	Y int
}

type Tetromino struct {
	Cells  []Point
	Letter rune
}
