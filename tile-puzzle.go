package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	board := [...]int{-1, -1, -1, -1, -1, -1, -1, -1, -1}
	initBoard(&board)
	printInstructions()

	reader := bufio.NewReader(os.Stdin)
	for {
		printBoard(&board)
		input, _ := reader.ReadString('\n')

		if input == "w\n" {
			up(&board)
		}
		if input == "a\n" {
			left(&board)
		}
		if input == "s\n" {
			down(&board)
		}
		if input == "d\n" {
			right(&board)
		}
	}

}

func initBoard(board *[9]int) {
	for i := 0; i < len(board); i++ {
		for {
			rand.Seed(time.Now().UTC().UnixNano())
			value := rand.Intn(len(board))
			if contains, _ := contains(board, value); !contains {
				board[i] = value
				break
			}
		}
	}
}

func printInstructions() {
	fmt.Println("Use w, a, s, d to move the space up, left, right and down")
}

func printBoard(board *[9]int) {
	for i, v := range board {
		if i%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", v)
	}
	fmt.Printf("\n")
}

func left(board *[9]int) bool {
	contains, pos := contains(board, 0)
	if !contains {
		panic("board do not contains 0.")
	}

	if pos%3 != 0 {
		board[pos] = board[pos-1]
		board[pos-1] = 0
		return true
	}
	return false
}

func right(board *[9]int) bool {
	contains, pos := contains(board, 0)
	if !contains {
		panic("board do not contains 0.")
	}

	if (pos <= 2 && pos%2 != 0) ||
		(pos >= 3 && pos <= 5 && pos%5 != 0) ||
		(pos >= 6 && pos%8 != 0) {
		board[pos] = board[pos+1]
		board[pos+1] = 0
		return true
	}
	return false
}

func up(board *[9]int) bool {
	contains, pos := contains(board, 0)
	if !contains {
		panic("board do not contains 0.")
	}

	if pos-3 >= 0 {
		board[pos] = board[pos-3]
		board[pos-3] = 0
		return true
	}
	return false
}

func down(board *[9]int) bool {
	contains, pos := contains(board, 0)
	if !contains {
		panic("board do not contains 0.")
	}

	if pos+3 < len(board) {
		board[pos] = board[pos+3]
		board[pos+3] = 0
		return true
	}
	return false
}

func contains(board *[9]int, e int) (bool, int) {
	for i, v := range board {
		if v == e {
			return true, i
		}
	}
	return false, -1
}
