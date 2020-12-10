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
Placing N queens causes the code to see if the current
board has already been encountered.
If not, it prints the board.
