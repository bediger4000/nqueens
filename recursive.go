package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

const (
	QUEEN  = -1
	EMPTY  = 0
	MARK   = 1
	UNMARK = -1
)

func main() {
	size := flag.Int("N", 5, "size of side of board")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()
	fmt.Printf("%d squares on a side\n", *size)

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var board [12][12]int

	checkBoard(0, *size, &board)

	fmt.Printf("%d unique %d-queens boards\n", uniqueBoardCount, *size)
}

var uniqueBoards = make(map[[144]byte]bool)
var uniqueBoardCount int

func stringify(sz int, board *[12][12]int) [144]byte {
	var buf [144]byte
	for i := 0; i < sz; i++ {
		for j := 0; j < sz; j++ {
			if (*board)[i][j] == QUEEN {
				buf[12*i+j] = 'Q'
			}
		}
	}
	return buf
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
	for i := 0; i < size; i++ {
		if i == q {
			continue
		}
		(*board)[p][i] += mark
	}
	// col with <p,q> in it
	for i := 0; i < size; i++ {
		if i == p {
			continue
		}
		(*board)[i][q] += mark
	}

	// diagonal, lower left to upper right
	for i := 1; p-i >= 0 && q-i >= 0; i++ {
		(*board)[p-i][q-i] += mark
	}
	for i := 1; p+i < size && q+i < size; i++ {
		(*board)[p+i][q+i] += mark
	}

	// diagonal, upper left to lower right
	for i := 1; p+i < size && q-i >= 0; i++ {
		(*board)[p+i][q-i] += mark
	}
	for i := 1; p-i >= 0 && q+i < size; i++ {
		(*board)[p-i][q+i] += mark
	}
}
