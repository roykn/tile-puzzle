package main

import "testing"

func TestContains(t *testing.T) {
	board := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	for i := range board {
		if contains, _ := contains(&board, i); !contains {
			t.Log(board)
			t.Errorf("board not contains %d", i)
		}
	}

	if contains, _ := contains(&board, 9); contains {
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

func TestLeft(t *testing.T) {
	tables := []struct {
		input  [9]int
		output [9]int
		result bool
	}{
		{[9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, [9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, true},
		{[9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, [9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, true},
		{[9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, [9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, false},
	}

	for _, table := range tables {
		result := left(&table.input)
		if (table.input != table.output) && (result != table.result) {
			t.Logf("left returned %t", result)
			t.Log(table.input)
			t.Log(table.output)
			t.Error("Move left failed.")
		}
	}
}

func TestRight(t *testing.T) {
	tables := []struct {
		input  [9]int
		output [9]int
		result bool
	}{
		{[9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, [9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, true},
		{[9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, true},
		{[9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, false},
	}

	for _, table := range tables {
		result := right(&table.input)
		if (table.input != table.output) && (result != table.result) {
			t.Logf("right returned %t", result)
			t.Log(table.input)
			t.Log(table.output)
			t.Error("Move right failed.")
		}
	}
}
