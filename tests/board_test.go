package main

import (
  "testing"
  "fmt"
)

type stringPos struct {
  s string
  p Position
}

func TestToPos(t *testing.T) {
  tests := []stringPos {
    stringPos {
      "a1", Position{0,0},
    },
    stringPos {
      "h8", Position{7,7},
    },
    stringPos {
      "d5", Position{4,3},
    },
  }

  for _, test := range tests {
    if ToPos(test.s) != test.p {
      t.Error(fmt.Sprintf("Expected %v, got %v", test.p, ToPos(test.s)))
    }
  }
}

type fenPosPiece struct {
  fen string
  pos Position
  piece Piece
}

func TestFindPiece(t *testing.T) {
  tests := []fenPosPiece {
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{0,0},
      &Rook{true},
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{7,7},
      &Rook{false},
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{1,0},
      &Pawn{true},
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{0,2},
      &Bishop{true},
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{6,1},
      &Pawn{false},
    },
  }
  for _, test := range(tests) {
    b := GenerateBoard(test.fen)
    found, piece := b.findPiece(test.pos)
    if !found || test.piece.ToString() != piece.ToString() {
        t.Error(fmt.Sprintf("Expected %v, got %v", test.piece.ToString(), piece.ToString()))
    }
  }
}

type fenCheck struct {
  fen string
  whiteCheck, blackCheck bool
}

func TestCheck(t *testing.T) {
  tests := []fenCheck{
    fenCheck{
      "rnbqk1nr/ppp2ppp/4p3/3p4/1b1P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1", true, false,
    },
    fenCheck{
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", false, false,
    },
    fenCheck{
      "rnbqk1nr/ppp3pp/4pp2/3p2B1/Qb1P4/2P5/PP2PPPP/RN2KBNR b KQkq - 0 1", false, true,
    },
  }
  for _, test := range(tests) {
    b := GenerateBoard(test.fen)
    if b.inCheck(true) != test.whiteCheck {
      t.Error(fmt.Sprintf("Expected white check to be %v, got %v", test.whiteCheck, b.inCheck(true)))
    }
    if b.inCheck(false) != test.blackCheck {
      t.Error(fmt.Sprintf("Expected black check to be %v, got %v", test.blackCheck, b.inCheck(false)))
    }
    if b.inCheck(false) && b.inCheck(true) {
      t.Error(fmt.Sprintf("Both players in check? Look at this position: %v", test.fen))
    }
  }
}
