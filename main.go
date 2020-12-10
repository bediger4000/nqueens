package main

import (
	"flag"
	"fmt"
)

func main() {
	size := flag.Int("N", 5, "size of side of board")
	flag.Parse()
	fmt.Printf("%d squares on a side\n", *size)

	board := make([][]int, *size)

	for i := 0; i < *size; i++ {
		board[i] = make([]int, *size)
	}

	for i := 0; i < *size; i++ {
		for j := 0; j < *size; j++ {
			setQueen(1, *size, &board, i, j)
		}
	}
}

func setQueen(n, size int, board *[][]int) {
	if n == size {
		printBoard(board)
		return
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (*board)[i][j] == 0 {
				// place queen
				(*board)[i][j] = -1
				// increment all this queen's accessible squares
				for x := 0; x < size; x++ {
					if x != i {
						(*board)[x][j]++
					}
				}
				for x := 0; x < size; x++ {
					if x != j {
						(*board)[i][x]++
					}
				}
				// call setQueen
				setQueen(n+1, size, board)
				// decrement all this queen's accessible squares
				// remove queen
				(*board)[i][j] = 0
			}
		}
	}
}

func printBoard(board *[][]int) {
	for _, row := range *board {
		for _, x := range row {
			marker := '_'
			if x == -1 {
				marker = 'Q'
			}
			fmt.Printf("%c", marker)
		}
		fmt.Println()
	}
}
