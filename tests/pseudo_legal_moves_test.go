package main

import (
  "testing"
  "fmt"
  // "reflect"
)

type TestCasePseudoMoves struct {
  fen string
  pos string
  expected []string
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func TestPseudoLegalMoves2(t *testing.T) {
  tests := []TestCasePseudoMoves {
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "d5", []string{ "d6",  "e5",  "d4",  "d3",  "c5",  "b5",  "a5",  "e6",  "f7",  "c4",  "c6",  "b7",},},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "c3", []string{ "a4",  "e2",  "b5",  "d1",  "b1",},},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "e4", []string{  },},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "b7", []string{ "b6",  "b5",  },},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "f1", []string{ "e2",  "d3",  "c4",  "b5",  "a6",},},
  }
  for _, test := range(tests) {
    b := GenerateBoard(test.fen)
    pos := ToPos(test.pos)
    _, piece := b.findPiece(pos)
    moves := piece.GetPseudoLegalMoves(pos, &b)
    sqs := []string{}
    for _, move := range(moves) {
      sqs = append(sqs, move.String())
    }
    for _, sq := range(sqs) {
      if !contains(test.expected, sq) {
        t.Error(fmt.Sprintf("Square %v not in Expected Range", sq))
      }
    }
    for _, sq := range(test.expected) {
      if !contains(sqs, sq) {
        t.Error(fmt.Sprintf("Square %v expected but not in output", sq))
      }
    }
  }
}

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
