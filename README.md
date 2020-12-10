# N Queens

Generalized [8 Queens problem](https://en.wikipedia.org/wiki/Eight_queens_puzzle)

## Building and Running

This code does not depend on any 3rd party packages.
I did not do go modules.

```sh
$ cd $GOPATH/src
$ git clone https://github.com/bediger4000/nqueens.git
$ cd nqueens
$ go build .
$ ./nqueens -N 8
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

## My Solution

[Code](main.go)

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
You could keep your own stack-of-positions around,
since only one board is in use at any given time.
