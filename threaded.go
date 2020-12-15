package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
)

const (
	QUEEN  = -1
	EMPTY  = 0
	MARK   = 1
	UNMARK = -1
)

func main() {
	numberThreads := runtime.NumCPU()
	size := flag.Int("N", 5, "size of side of board")
	threads := flag.Int("t", numberThreads, "number of threads")
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

	numberThreads = *threads
	runtime.GOMAXPROCS(numberThreads)

	jobs, feedback := startJobs(*size, numberThreads)

	go func() {
		jobcount := 0
		for m := 0; m < *size; m++ {
			for n := 0; n < *size; n++ {
				jobs <- &job{size: *size, i: m, j: n}
				jobcount++
			}
		}
	}()

	done := 0
	for _ = range feedback {
		done++
		if done == *size**size {
			break
		}
	}

	fmt.Printf("%d unique %d-queens boards\n", uniqueBoardCount, *size)
}

func startJobs(size int, numberThreads int) (chan *job, chan int) {
	jobs := make(chan *job, 0)
	done := make(chan int, numberThreads)

	for i := 0; i < numberThreads; i++ {
		go doJobs(i, jobs, done)
	}

	return jobs, done
}

func doJobs(serialNumber int, jobs chan *job, done chan int) {
	var board [12][12]int
	for j := range jobs {
		board[j.i][j.j] = QUEEN
		markSquares(j.size, &board, j.i, j.j, MARK)
		checkBoard(1, j.size, &board)
		markSquares(j.size, &board, j.i, j.j, UNMARK)
		board[j.i][j.j] = EMPTY
		done <- serialNumber
	}
}

type job struct {
	size int
	i    int
	j    int
}

var uniqueLock sync.Mutex
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

func collectReports(sz int, board *[12][12]int) {
	uniqueLock.Lock()
	boardAsString := stringify(sz, board)
	if !uniqueBoards[boardAsString] {
		printBoard(sz, board)
		uniqueBoards[boardAsString] = true
		uniqueBoardCount++
	}
	uniqueLock.Unlock()
}

func checkBoard(ply, size int, board *[12][12]int) {
	if ply == size {
		// All queens placed, base recursion case
		collectReports(size, board)
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
