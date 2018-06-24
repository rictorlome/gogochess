package main

import (
  "testing"
  "fmt"
  // "reflect"
)

func TestPseudoLegalMoves(t *testing.T) {
  var fenOne string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/1NBQKBNR w KQkq - 0 1"
  b := GenerateBoard(fenOne)

  r, bish, king, queen := Rook{}, Bishop{}, King{}, Queen{}
  r.isWhite, bish.isWhite, king.isWhite, queen.isWhite = true, true, true, true

  if m := r.GetPseudoLegalMoves(Position{0,0}, &b); len(m) > 0 {
    t.Error(fmt.Sprintf("Expected trapped %T at A1 to have 0 moves, got %d", r, len(m)))
  }
  if m := bish.GetPseudoLegalMoves(Position{0,0}, &b); len(m) > 0 {
    t.Error(fmt.Sprintf("Expected trapped %T at A1 to have 0 moves, got %d", bish, len(m)))
  }
  if m := king.GetPseudoLegalMoves(Position{0,0}, &b); len(m) > 0 {
    t.Error(fmt.Sprintf("Expected trapped %T at A1 to have 0 moves, got %d", king, len(m)))
  }
  if m := queen.GetPseudoLegalMoves(Position{0,0}, &b); len(m) > 0 {
    t.Error(fmt.Sprintf("Expected trapped %T at A1 to have 0 moves, got %d", queen, len(m)))
  }
}
