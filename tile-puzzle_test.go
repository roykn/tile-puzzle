package main

import "testing"

func TestContains(t *testing.T) {
	board := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	for i := range board {
		if !contains(&board, i) {
			t.Log(board)
			t.Errorf("board not contains %d", i)
		}
	}

	if contains(&board, 9) {
		t.Errorf("board should not contain %d", 9)
	}
}

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
