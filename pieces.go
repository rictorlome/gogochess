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

// var whitePieces = map[Position]Piece
// var blackPieces = map[Position]Piece

func main() {
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  fields := strings.Split(fen, " ")
  pos := fields[0]
  rows := strings.Split(pos, "/")
  for _, row := range rows {
    // offset := 0
    for _, sq := range row {
      _, found := whiteSymbols[string(sq)]
      fmt.Println(found)
    }
  }

  // p := Position{1,1}
  // b := Board{}
  // b.grid[0] = &p
  // fmt.Println(p,b)
}
