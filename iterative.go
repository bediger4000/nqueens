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

type position struct {
	i int
	j int
}

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

	sz := *size
	var board [12][12]int

	var stack [12]position
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

				stack[queencount].i = i
				stack[queencount].j = j
				queencount++

				i, j = 0, 0

				if queencount == sz {
					// Got to N queens, backtrack to N-1 queens
					printUniqueBoards(sz, &board)
					break OUT
				}
			}
			j = 0
		}

		// loop terminating condition: no queen on the board.
		// This should happen after the last of NxN squares on
		// the board gets the first queen. When the last square
		// has its (first layer) queen popped, <i, j> set to <N-1, N>
		// by the "remove top queen" code below.
		// The for-loop over i has one final iteration, but since j == N,
		// the inner for-loop-over-j doesn't do any work. It just jumps
		// over the code. No queen position gets pushed on the stack-of-positions.
		if queencount == 0 {
			break
		}

		// pop top queen's position off the stack
		queencount--

		// remove top queen from board
		i, j = stack[queencount].i, stack[queencount].j
		markSquares(sz, &board, i, j, UNMARK)
		board[i][j] = EMPTY

		// set j to be one square to the right of queen's position.
		j++
	}

	fmt.Printf("%d unique %d-queens boards\n", uniqueBoardCount, sz)
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
