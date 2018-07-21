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

func TestPseudoLegalMoves(t *testing.T) {
  tests := []TestCasePseudoMoves {
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "d5", []string{ "d6",  "e5",  "d4",  "d3",  "c5",  "b5",  "a5",  "e6",  "f7",  "c4",  "c6",  "b7"}},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "c3", []string{ "a4",  "e2",  "b5",  "d1",  "b1"}},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "e4", []string{  }},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "b7", []string{ "b6",  "b5"  }},
    TestCasePseudoMoves{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "f1", []string{ "e2",  "d3",  "c4",  "b5",  "a6"}},
    TestCasePseudoMoves{"rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 1", "e5", []string{ "f6",  "e6",  }},
    TestCasePseudoMoves{"rnb1kb1r/ppp1p1pp/3q1n2/3pPp2/3P4/8/PPP2PPP/RNBQKBNR w KQkq - 0 1", "e5", []string{ "d6",  "f6",  "e6",  }},
    // TestCasePseudoMoves{"rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 1", "e5", []string{ "e2",  "d3",  "c4",  "b5",  "a6"}},
  }
  for _, test := range(tests) {
    b := GenerateBoard(test.fen)
    pos := ToPos(test.pos)
    _, piece := b.findPiece(pos)
    moves := piece.GetPseudoLegalMoves(pos, &b)
    sqs := []string{}
    for move, _ := range(moves) {
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
