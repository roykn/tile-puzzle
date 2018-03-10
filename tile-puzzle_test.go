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

func TestMove(t *testing.T) {
	tables := []struct {
		input  [9]int
		output [9]int
		result bool
		fn     func(*[9]int) bool
		fnName string
	}{
		{[9]int{0, 1, 2, 3, 4, 5, 7, 8, 6}, [9]int{1, 0, 2, 3, 4, 5, 7, 8, 6}, true, left, "right"},
		{[9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, true, left, "left"},
		{[9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, [9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, true, left, "left"},
		{[9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, [9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, false, left, "left"},
		{[9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, [9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, true, right, "right"},
		{[9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, true, right, "right"},
		{[9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, false, right, "right"},
		{[9]int{1, 2, 0, 3, 4, 5, 7, 8, 6}, [9]int{1, 2, 0, 3, 4, 5, 7, 8, 6}, false, right, "right"},
		{[9]int{1, 2, 3, 4, 5, 6, 7, 0, 8}, [9]int{1, 2, 3, 4, 0, 6, 7, 5, 8}, true, up, "up"},
		{[9]int{1, 2, 3, 4, 0, 6, 7, 5, 8}, [9]int{1, 0, 3, 4, 2, 6, 7, 5, 8}, true, up, "up"},
		{[9]int{1, 0, 3, 4, 2, 6, 7, 5, 8}, [9]int{1, 0, 3, 4, 2, 6, 7, 5, 8}, false, up, "up"},
		{[9]int{1, 0, 2, 3, 4, 5, 6, 7, 8}, [9]int{1, 4, 2, 3, 0, 5, 6, 7, 8}, true, down, "down"},
		{[9]int{1, 4, 2, 3, 0, 5, 6, 7, 8}, [9]int{1, 4, 2, 3, 7, 5, 6, 0, 8}, true, down, "down"},
		{[9]int{1, 4, 2, 3, 7, 5, 6, 0, 8}, [9]int{1, 4, 2, 3, 7, 5, 6, 0, 8}, false, down, "down"},
	}

	for _, table := range tables {
		result := table.fn(&table.input)
		if result != table.result {
			t.Logf("%s returned %t", table.fnName, result)
			t.Logf("output:          %v", table.input)
			t.Logf("expected output: %v", table.output)
			t.Errorf("Move %s failed", table.fnName)
		}
	}
}
