package main

import (
  "strings"
  "strconv"
)

func GenerateBoard(fen string) Board {
  b := Board{}
  fields := strings.Split(fen, " ")

  b.whites, b.blacks, b.whiteKing, b.blackKing = GeneratePositions(fields[0])
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

func GeneratePositions(pos string) (map[Position]Piece, map[Position]Piece, Position, Position) {
    rows := strings.Split(pos, "/")
    whites := make(map[Position]Piece)
    blacks := make(map[Position]Piece)
    var whiteKing, blackKing Position
    for i, row := range rows {
      offset := 0
      for j, sq := range row {
        piece := ToPiece(string(sq))
        if piece == nil {
          // This gives the numeric value of the rune(ie '56' to the int 8)
          // Minus one because the offset number takes up a square itself
          offset += int(sq - '0') - 1
        } else {
          pos := Position{7-i,j+offset}
          if piece.IsWhite() {
            whites[pos] = piece
            if piece.ToString() == "K" {
              whiteKing = pos
            }
          } else  {
            blacks[pos] = piece
            if piece.ToString() == "k" {
              blackKing = pos
            }
          }
        }
      }
    }
    return whites, blacks, whiteKing, blackKing
}
