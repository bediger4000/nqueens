# N Queens

Generalized [8 Queens problem](https://en.wikipedia.org/wiki/Eight_queens_puzzle)

## Building and Running

This code does not depend on any 3rd party packages.
I did not do go modules.

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
runs and outputs exactly the same way.

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

This solution does backtracking via stack of positions
of where it last placed a queen.
This was a lot harder than I thought it would be.
I wanted to place every queen, even the first one,
with the same code.
I found this to be impossible.
The iterative backtracking is hard to get to terminate if you
only have x,y position on board and a stack depth to guide you.

The iterative version takes about twice as long as the recursive
version for a given number of queens (and hence, board size).
