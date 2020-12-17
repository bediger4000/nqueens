# N Queens

Generalized [8 Queens problem](https://en.wikipedia.org/wiki/Eight_queens_puzzle)

While not a "Daily Coding Problem",
I've seen a 5-queens solution posed as a programming job interview question.
I think I'd rate it at least "medium" on the Daily Coding Problem scale.
The iterative version might even be a "hard".

The idea is to place N queen chess pieces on an N by N chessboard
so that no queen can capture any other queen.
That is, no two queens can share a row, a column or a diagonal.

## Building and Running

I wrote three solutions in [Go](https://golang.org/).

This code does not depend on any 3rd party packages.
I did not do Go modules.

To build the [recursive](recursive.go) version:

```sh
$ cd $GOPATH/src
$ git clone https://github.com/bediger4000/nqueens.git
$ cd nqueens
$ go build recursive.go
$ ./recursive -N 8
8 squares on a side
Q.......
....Q...
.......Q
.....Q..
..Q.....
......Q.
.Q......
...Q....

   ...
92 unique 8-queens boards
$
```

There's also an [iterative](iterative.go) version that builds,
runs and outputs exactly the same way,
except it's named `iterative.go`.

I wrote a [threaded](threaded.go) version of the recursive algorithm.
It starts one goroutine per square on the board,
each goroutine finds all the N-queen solutions it can.
Only unique solutions are reported.

Both versions can handle a maximum of a 12x12 board.
You will die of old age waiting for all 12x12 board's solutions.

## My Solutions

* [Recursive version](recursive.go)

I implemented a recursive backtracking solution.
Placing a queen increments a count for every square in
that queen's row, column and diagonals.
Queens can only be placed on zero-count squares.
Once the function places a queen,
it calls iself to place the next queen.
Placing N queens causes the code to see if the current
board has already been encountered.
If not, it prints the board.
This function, in Go, uses a pointer to a board of type `[][]int`
to avoid copying a bunch of slices all over the place.

The recursive function mostly exists to keep track of where
on the chess board a queen got placed.

* [Threaded version](threaded.go)

Keeps a pool of goroutines blocked on a channel,
sending `(i,j)` coordinates of the first queen to place on
an oterwise empty board.
Each goroutine finds all the N-queen placements resulting from
recursive backtracking from that starting queen.
Only unique configurations are counted and reported.
I mixed CSP-style concurrency (goroutine blocked on a channel)
and the usual mutex-lock concurrency,
which I used when filtering and counting unique configurations.
The code uses a `*[12][12]int` pointer-to-array as the chess board
representation, to avoid hidden array copying when recursing.
That means that a successful configuration is a pointer,
which would cause problems if passed through a channel to a goroutine
managing filtering and counting.

* [Iterative version](iterative.go)

This solution does backtracking via an explicit stack of positions of where it
last placed a queen.
This was a lot harder than I thought it would be.
I wanted to place every queen,
even the first one,
with the same code.
In my first attempt, I tried to get clever by 
not treating the Nth-queen-placement exactly the same
as all other queen's placement.
This caused me to have a lot of difficulty terminating
the looping and getting the first queen placed on all squares of the board.
I finally got this to work by eliminating the optimizing of the special
case of the first queen's placement.

### Performance Comparison

I got iterative and recursive versions to have
roughly the same performance by eliminating all slice accesses,
and all dynamic allocations.
The stack of positions and the board become fixed-size arrays.
In the iterative version,
I used a discrete variable to track stack depth,
so I didn't need to use len(stack) anywhere.
The stack is a fixed size array of `struct position`, so no dynamic allocation
or garbage collection takes place.

This does limit the maximum board size.
I used 12x12, because I lost patience at running an 8x8 test case.

In a way, this performance enhancement seems to confirm the
1994 paper [Garbage Collection is Fast, But a Stack is Faster](http://dspace.mit.edu/handle/1721.1/6622).
Their "stack" is the call-frame-stack managed by hardware to
do stack-discipline function calls.
My two versions have a hardware stack (recursive version)
and a software stack (iterative version).
I had to eliminate dynamic allocation to get performance.
The only question is whether indirect accesses
(which are somewhat hidden from the programmer by Go's syntax
and runtime) were the performance problem,
or if it was dynamic allocation and garbage collection.

For once, the threaded version can outperform the single-threaded algorithms.
I did have to profile and optimize to get there, however.

#### Performance Testing

* Dell PowerEdge R530, runtime.NumCPU reports 40
* Dell Latitude E6420, runtime.NumCPU reports 4

I used 8 queens, which finds 92 unique solutions.

* Single-threaded

|machine|recursive|iterative|
|-------|------|------|
|R530   |17.7  |7.3 |
|E6420  |12.8  |8.8 |

Single-threaded versions give about the same numbers on both machines.

* Multi-threaded

|Threads|R530|E6420|
|-------|----|------|
|1|7.689|9.314|
|2|4.061|5.692|
|5|1.839|4.989|
|10|1.247|4.624|
|15|1.514|4.596|
|20|1.699|4.555|
|25|1.835|4.687|
|30|2.225|4.566|
|35|2.126|4.581|
|40|2.408|4.696|

The multi-threaded version is just a number of recursive versions in parallel,
so it's not too surprising that 1 thread takes about as long as the
single-threaded version.
I'm amused that the R530 bottoms out at about 1.25 seconds, and 10 threads,
while the E6420 bottoms out at 4.6 seconds, 15 threads.
I suspect that the variation at 5 threads and above is just noise.

Since the R530 has more CPUs and more hyperthreads,
it's not too surprising that it out-performs the E6420.
What I don't understand is why the R530 bottoms out at 10 threads,
given the large number of CPUs available.

#### Performance Tips

* I would not have improved performance without Go's [pprof]().
Profile first, even though sometimes it just shows your
program spending a lot of time in `syscall.Syscall` sometimes.
* Accessing a slice-of-slices is somewhat slower than
accessing a 2-D array.
* Creating a string from a slice-of-slices can be very slow.
* Using a smaller (144) array of bytes as a map key
is seemingly faster than using a string.
* A single for-loop with an if-statement in it to avoid
a particular index is slower than 2 for-loops whose indices
skip that particular index.
* Lots of dynamic allocation (i.e. `p := &position{i:x, j:y}`
causes slowdowns.
I'm not sure if dynamic allocation or garbage collection
takes the time.
