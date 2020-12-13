package main

import (
	"bytes"
	"flag"
	"fmt"
)

const (
	QUEEN  = -1
	EMPTY  = 0
	MARK   = 1
	UNMARK = -1
)

type position struct {
	i int
	j int
}

func main() {
	size := flag.Int("N", 5, "size of side of board")
	flag.Parse()
	fmt.Printf("%d squares on a side\n", *size)

	var board [][]int
	sz := *size

	for i := 0; i < sz; i++ {
		board = append(board, make([]int, sz))
	}

	var stack []*position
	var queencount int

	i, j := 0, 0
	for {
	OUT:
		for ; i < sz; i++ {
			for ; j < sz; j++ {
				if board[i][j] != EMPTY {
					continue
				}
				board[i][j] = QUEEN
				markSquares(sz, &board, i, j, MARK)
				queencount++

				stack = append(stack, &position{i: i, j: j})
				i, j = 0, 0

				if queencount == sz {
					printUniqueBoards(&board)
					break OUT
				}
			}
			j = 0
		}

		if queencount == 0 {
			break
		}

		// pop last queen's position off the stack
		queencount--
		pos := stack[queencount]
		stack = stack[:queencount]

		// remove that queen
		i, j = pos.i, pos.j
		markSquares(sz, &board, i, j, UNMARK)
		board[i][j] = EMPTY

		// set j to be one square to the right of queen's position.
		j++
	}

	fmt.Printf("%d unique %d-queens boards\n", uniqueBoardCount, sz)
}

var uniqueBoards = make(map[string]bool)
var uniqueBoardCount int

func stringify(board *[][]int) string {
	buf := bytes.Buffer{}
	for _, row := range *board {
		for _, x := range row {
			mark := byte('.')
			if x == QUEEN {
				mark = byte('Q')
			}
			buf.WriteByte(mark)
		}
	}
	return buf.String()
}

func printUniqueBoards(board *[][]int) {
	boardAsString := stringify(board)
	if !uniqueBoards[boardAsString] {
		printBoard(board)
		uniqueBoards[boardAsString] = true
		uniqueBoardCount++
	}
}

func printBoard(board *[][]int) {
	for _, row := range *board {
		for _, x := range row {
			marker := '.'
			if x == QUEEN {
				marker = 'Q'
			}
			fmt.Printf("%c", marker)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func printRawBoard(board *[][]int) {
	for _, row := range *board {
		for _, x := range row {
			fmt.Printf("%2d", x)
		}
		fmt.Println()
	}
	fmt.Println()
}

func markSquares(size int, board *[][]int, p, q, mark int) {
	// row with <p,q> in it
	for i := -size; i < size; i++ {
		if i == 0 {
			continue
		}
		n := q + i
		if n < 0 {
			continue
		}
		if n >= size {
			continue
		}
		(*board)[p][n] += mark
	}
	// col with <p,q> in it
	for i := -size; i < size; i++ {
		if i == 0 {
			continue
		}
		m := p + i
		if m < 0 {
			continue
		}
		if m >= size {
			continue
		}
		(*board)[m][q] += mark
	}

	// diagonal, lower left to upper right
	for i := -size; i < size; i++ {
		if i == 0 {
			continue
		}
		m := p + i
		if m < 0 {
			continue
		}
		if m >= size {
			continue
		}
		n := q + i
		if n < 0 {
			continue
		}
		if n >= size {
			continue
		}
		(*board)[m][n] += mark
	}
	// diagonal, upper left to lower right
	for i := -size; i < size; i++ {
		if i == 0 {
			continue
		}
		m := p - i
		if m < 0 {
			continue
		}
		if m >= size {
			continue
		}
		n := q + i
		if n < 0 {
			continue
		}
		if n >= size {
			continue
		}
		(*board)[m][n] += mark
	}
}
