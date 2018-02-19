package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var board [9]int
	initBoard(&board)
	printBoard(&board)
}

func initBoard(board *[9]int) {
	for i := 0; i < len(board)-1; i++ {
		for {
			rand.Seed(time.Now().UTC().UnixNano())
			value := rand.Intn(len(board))
			if !contains(board, value) {
				board[i] = value
				break
			}
		}
	}
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
	e := errors.New("Not implemented")
	panic(e)
}

func right(board *[9]int) bool {
	e := errors.New("Not implemented")
	panic(e)
}

func up(board *[9]int) bool {
	e := errors.New("Not implemented")
	panic(e)
}

func down(board *[9]int) bool {
	e := errors.New("Not implemented")
	panic(e)
}

func contains(board *[9]int, e int) bool {
	for _, v := range board {
		if v == e {
			return true
		}
	}
	return false
}
