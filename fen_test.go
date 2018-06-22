package main

import (
  "testing"
)

func TestGenerateBoard(t *testing.T) {
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  b := GenerateBoard(fen)
  // Temporary tests to check positions are accurate
  for pos, piece := range(b.whites) {
    if piece == string("P") {
      if pos.row != 1 {
        t.Error("Expected white pawn on 1st row, got ", pos.row)
      }
    } else {
      if pos.row != 0 {
        t.Error("Expected white piece on 0th row, got ", pos.row)
      }
    }
  }
  for pos, piece := range(b.blacks) {
    if piece == string("p") {
      if pos.row != 6 {
        t.Error("Expected white pawn on 6th row, got ", pos.row)
      }
    } else {
      if pos.row != 7 {
        t.Error("Expected white piece on 7th row, got ", pos.row)
      }
    }
  }
  if !b.whiteToMove {
    t.Error("Expected true, got ", b.whiteToMove)
  }
  if len(b.availableCastles) != 4 {
    t.Error("Expected 4 available castles, got ", len(b.availableCastles))
  }
  if b.enPassantSquare.row != -1 || b.enPassantSquare.col != -1 {
    t.Error("Expected null position, got ", b.enPassantSquare)
  }
}
