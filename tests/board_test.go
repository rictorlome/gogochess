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
  piece string
}

func TestFindPiece(t *testing.T) {
  tests := []fenPosPiece {
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{0,0},
      "R",
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{7,7},
      "r",
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{1,0},
      "P",
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{0,2},
      "B",
    },
    fenPosPiece {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      Position{6,1},
      "p",
    },
  }
  for _, test := range(tests) {
    b := GenerateBoard(test.fen)
    if test.piece != b.findPiece(test.pos) {
        t.Error(fmt.Sprintf("Expected %v, got %v", test.piece, b.findPiece(test.pos)))
    }
  }
}
