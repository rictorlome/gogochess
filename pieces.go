package main

import (
  "fmt"
  "strings"
  "time"
)

/*
rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
*/

type Position struct {
  row, col int
}

type Board struct {
  whites, blacks map[Position]string
  whiteToMove bool
  availableCastles map[string]bool
  enPassantSquare Position
}

type Piece interface {
  GetMoves(p Position, f Board) []Position
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

func GenerateBoard(fen string) Board {
  b := Board{}
  fields := strings.Split(fen, " ")

  b.whites, b.blacks = GeneratePositions(fields[0])
  b.whiteToMove = (fields[1] == string('w'))
  b.availableCastles = SetAvailableCastles(fields[2])
  b.enPassantSquare = SetEnpassantSquare(fields[3])

  return b
}

func SetAvailableCastles(avail string) map[string]bool {
  availableCastles := map[string]bool{"bk" : false, "bq": false, "wk": false, "wq": false}
  if strings.Contains(avail, "K") {
    availableCastles["wk"] = true
  }
  if strings.Contains(avail, "Q") {
    availableCastles["wq"] = true
  }
  if strings.Contains(avail, "k") {
    availableCastles["bk"] = true
  }
  if strings.Contains(avail, "q") {
    availableCastles["bq"] = true
  }
  return availableCastles
}

func SetEnpassantSquare(algebriac string) Position {
    return Position{-1,-1}
}

func GeneratePositions(pos string) (map[Position]string, map[Position]string) {
    rows := strings.Split(pos, "/")
    whites := make(map[Position]string)
    blacks := make(map[Position]string)

    for i, row := range rows {
      offset := 0
      for j, sq := range row {
        _, white := whiteSymbols[string(sq)]
        _, black := blackSymbols[string(sq)]
        switch {
        case white:
          whites[Position{7-i,j+offset}] = string(sq)
        case black:
          blacks[Position{7-i,j+offset}] = string(sq)
        default:
          // This gives the numeric value of the rune(ie '56' to the int 8)
          offset += int(sq - '0')
        }
      }
    }
    return whites, blacks
}

func main() {
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  start := time.Now()
  b := GenerateBoard(fen)

  fmt.Println(b.whites)
  fmt.Println(b.blacks)
  fmt.Println(b.whiteToMove)
  fmt.Println(b.availableCastles)
  fmt.Println(b.enPassantSquare)
  t := time.Now()
  fmt.Println(t.Sub(start))
}
