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

	for i := 0; i < *size; i++ {
		board = append(board, make([]int, *size))
	}

	var stack []*position
	var queencount int
	for m := 0; m < *size; m++ {
		for n := 0; n < *size; n++ {
			board[m][n] = QUEEN
			markSquares(*size, &board, m, n, MARK)
			queencount++

			i, j := 0, 0
			for {
			OUT:
				for ; i < *size; i++ {
					for ; j < *size; j++ {
						if board[i][j] != EMPTY {
							continue
						}
						board[i][j] = QUEEN
						markSquares(*size, &board, i, j, MARK)
						queencount++

						stack = append(stack, &position{i: i, j: j})
						i, j = 0, 0

						if queencount == *size {
							printUniqueBoards(&board)
							break OUT
						}
					}
					j = 0
				}
				if len(stack) == 0 {
					break
				}
				l := len(stack) - 1
				pos := stack[l]
				stack = stack[:l]
				i, j = pos.i, pos.j
				markSquares(*size, &board, i, j, UNMARK)
				board[i][j] = EMPTY
				queencount--
				j++
			}

			board[m][n] = EMPTY
			markSquares(*size, &board, m, n, UNMARK)
			queencount--
		}
	}

	fmt.Printf("%d unique %d-queens boards\n", uniqueBoardCount, *size)
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