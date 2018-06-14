package main

import (
  "fmt"
  "strings"
)

/*
rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
*/

type Position struct {
  row, col int
}

type Piece struct {
  symbol string
}

var whiteSymbols = map[string]bool{
  "R": true,
  "N": true,
  "B": true,
  "Q": true,
  "K": true,
  "P": true,
}
var blackSymbols = map[string]bool{
  "r": true,
  "n": true,
  "b": true,
  "q": true,
  "k": true,
  "p": true,
}

var whites map[Position]string
var blacks map[Position]string

func main() {
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  fields := strings.Split(fen, " ")
  pos := fields[0]
  rows := strings.Split(pos, "/")

  whites = make(map[Position]string)
  blacks = make(map[Position]string)

  for i, row := range rows {
    offset := 0
    for j, sq := range row {
      _, white := whiteSymbols[string(sq)]
      _, black := blackSymbols[string(sq)]
      if white {
        whites[Position{7-i,j+offset}] = string(sq)
        continue
      }
      if black {
        blacks[Position{7-i,j+offset}] = string(sq)
        continue
      }
      // This gives the numeric value of the rune(ie '56' to the int 8)
      offset += int(sq - '0')
    }
  }
  fmt.Println(whites)
  fmt.Println(blacks)

  // p := Position{1,1}
  // b := Board{}
  // b.grid[0] = &p
  // fmt.Println(p,b)
}
