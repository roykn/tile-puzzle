package main

import (
	"fmt"
)

func main() {
	var board [9]int = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 0}
	initBoard(&board)
}

func initBoard(board *[9]int) {
	for i, v := range board {
		fmt.Printf("%d %d\n", i, v)
	}
}

func printBoard(board *[9]int) {

}

func left(board *[9]int) {

}

func right(board *[9]int) {

}

func up(board *[9]int) {

}

func down(board *[9]int) {

}
