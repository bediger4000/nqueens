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

func main() {
	size := flag.Int("N", 5, "size of side of board")
	flag.Parse()
	fmt.Printf("%d squares on a side\n", *size)

	var board [12][12]int

	checkBoard(0, *size, &board)

	fmt.Printf("%d unique %d-queens boards\n", uniqueBoardCount, *size)
}

var uniqueBoards = make(map[string]bool)
var uniqueBoardCount int

func stringify(sz int, board *[12][12]int) string {
	buf := bytes.Buffer{}
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			mark := byte('.')
			if (*board)[i][j] == QUEEN {
				mark = byte('Q')
			}
			buf.WriteByte(mark)
		}
	}
	return buf.String()
}

func printUniqueBoards(sz int, board *[12][12]int) {
	boardAsString := stringify(sz, board)
	if !uniqueBoards[boardAsString] {
		printBoard(sz, board)
		uniqueBoards[boardAsString] = true
		uniqueBoardCount++
	}
}

func checkBoard(ply, size int, board *[12][12]int) {
	if ply == size {
		// All queens placed, base recursion case
		printUniqueBoards(size, board)
		return
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (*board)[i][j] != EMPTY {
				continue
			}
			(*board)[i][j] = QUEEN
			markSquares(size, board, i, j, MARK)
			checkBoard(ply+1, size, board)
			markSquares(size, board, i, j, UNMARK)
			(*board)[i][j] = EMPTY
		}
	}
}

func printBoard(sz int, board *[12][12]int) {
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			marker := '.'
			if (*board)[i][j] == QUEEN {
				marker = 'Q'
			}
			fmt.Printf("%c", marker)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}

func markSquares(size int, board *[12][12]int, p, q, mark int) {
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
