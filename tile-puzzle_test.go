package main

import "testing"

func TestInitBoard(t *testing.T) {
	var board [9]int
	initBoard(&board)
	t.Log(board)

	// every board value should be unique
	for i, value := range board {
		for j, other := range board {
			if value == other && i != j {
				t.Errorf("%d is not unique.", value)
			}
		}
	}
}
