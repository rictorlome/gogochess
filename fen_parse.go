package main

import (
  "strings"
  "strconv"
)

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
  b.whiteToMove = (fields[1] == "w")
  b.availableCastles = SetAvailableCastles(fields[2])
  b.enPassantSquare = SetEnpassantSquare(fields[3])

  halfMoveClock, err := strconv.Atoi(fields[4])
  if err == nil {
    b.halfMoveClock = halfMoveClock
  }
  fullMoveNumber, err := strconv.Atoi(fields[5])
  if err == nil {
    b.fullMoveNumber = fullMoveNumber
  }

  return b
}

func SetAvailableCastles(avail string) map[string]bool {
  return map[string]bool {
     "bk": strings.Contains(avail, "k"),
     "bq": strings.Contains(avail, "q"),
     "wk": strings.Contains(avail, "K"),
     "wq": strings.Contains(avail, "Q"),
  }
}

func SetEnpassantSquare(algebraic string) Position {
  row, col := -1, -1
  if len(algebraic) == 2 {
    cols := "abcdefgh"
    row = int(algebraic[1] - '0') - 1
    col = strings.Index(cols, string(algebraic[0]))
  }
  return Position{row,col}
}

func GeneratePositions(pos string) (map[Position]Piece, map[Position]Piece) {
    rows := strings.Split(pos, "/")
    whites := make(map[Position]Piece)
    blacks := make(map[Position]Piece)

    for i, row := range rows {
      offset := 0
      for j, sq := range row {
        piece := ToPiece(string(sq))
        if piece == nil {
          // This gives the numeric value of the rune(ie '56' to the int 8)
          // Minus one because the offset number takes up a square itself
          offset += int(sq - '0') - 1
        } else {
          if piece.IsWhite() {
            whites[Position{7-i,j+offset}] = piece
          } else  {
            blacks[Position{7-i,j+offset}] = piece
          }
        }
      }
    }
    return whites, blacks
}
