package main

import (
	"testing"
)

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
	board := [...]int{-1, -1, -1, -1, -1, -1, -1, -1, -1}
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

func TestSolved(t *testing.T) {
	tables := []struct {
		input  [9]int
		result bool
		fn     func(*[9]int) bool
		fnName string
	}{
		{[9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, true, solved, "solved"},
		{[9]int{0, 1, 2, 3, 5, 4, 6, 7, 8}, false, solved, "solved"},
	}

	for _, table := range tables {
		result := table.fn(&table.input)
		if result != table.result {
			t.Logf("%s returned %t", table.fnName, result)
			t.Logf("expected output: %v", table.result)
			t.Errorf("%v failed", table.fnName)
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
		{[9]int{0, 1, 2, 3, 4, 5, 7, 8, 6}, [9]int{1, 0, 2, 3, 4, 5, 7, 8, 6}, true, left, "1 moves left"},
		{[9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, true, left, "5 moves left"},
		{[9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, [9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, true, left, "4 moves left"},
		{[9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, [9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, false, left, "nothing moves left"},
		{[9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, [9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, true, right, "4 moves right"},
		{[9]int{1, 2, 3, 4, 5, 0, 7, 8, 6}, [9]int{1, 2, 3, 4, 0, 5, 7, 8, 6}, true, right, "5 moves right"},
		{[9]int{1, 2, 0, 3, 4, 5, 7, 8, 6}, [9]int{1, 0, 2, 3, 4, 5, 7, 8, 6}, true, right, "2 moves right"},
		{[9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, [9]int{1, 2, 3, 0, 4, 5, 7, 8, 6}, false, right, "nothing moves right"},
		{[9]int{9, 9, 9, 9, 0, 9, 9, 5, 9}, [9]int{9, 9, 9, 9, 5, 9, 9, 0, 9}, true, up, "5 moves up"},
		{[9]int{9, 0, 9, 9, 2, 9, 9, 9, 9}, [9]int{9, 2, 9, 9, 0, 9, 9, 9, 9}, true, up, "2 moves up"},
		{[9]int{1, 2, 3, 4, 5, 6, 7, 0, 8}, [9]int{1, 2, 3, 4, 5, 6, 7, 0, 8}, false, up, "nothing moves up"},
		/*{[9]int{1, 0, 2, 3, 4, 5, 6, 7, 8}, [9]int{1, 0, 2, 3, 4, 5, 6, 7, 8}, false, down, "nothing moves down"},
		{[9]int{1, 4, 2, 3, 0, 5, 6, 7, 8}, [9]int{1, 0, 2, 3, 4, 5, 6, 7, 8}, true, down, "4 moves down"},
		{[9]int{1, 4, 2, 3, 7, 5, 6, 0, 8}, [9]int{1, 4, 2, 3, 0, 5, 6, 7, 8}, true, down, "7 moves down"},*/
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
