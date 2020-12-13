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

I wrote two solutions in [Go](https://golang.org/).

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
