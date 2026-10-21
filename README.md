# Tetris Optimizer

A Go program that reads tetrominoes from a text file and arranges them inside the smallest possible square.

Each tetromino is represented by an uppercase letter. The first piece is `A`, the second is `B`, and so on. Empty cells are represented by `.`.

## Input Format

Each tetromino must:

- Contain exactly four lines
- Have exactly four characters per line
- Use only `#` and `.`
- Contain exactly four connected `#` cells
- Be separated from the next tetromino by one empty line

Example:

```text
#...
#...
#...
#...

....
....
..##
..##