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

type TestMoveStruct struct {
  uci string
  fen string
}
// Generated using lichess.org/analysis
// The only difference is: I add enpassant sq to state after every 2-square-advance, regardless of where opposing pawns are.
// PGN below:
/*
1. e4 e5 2. Nf3 Nc6 3. Bb5 Nf6 4. O-O d6
5. Re1 a6 6. Ba4 b5 7. Bb3 b4 8. c4 bxc3
9. Nxc3 Be7 10. d4 O-O 11. dxe5 Nd4 12. exf6 Nxb3
13. fxe7 Nxc1 14. exd8=Q Kh8 15. Qxf8#
*/
func TestMove(t *testing.T) {
  tests := []TestMoveStruct {
    TestMoveStruct{
      "e2e4", "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1",
    },
    TestMoveStruct{
      "e7e5", "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2",
    },
    TestMoveStruct{
      "g1f3", "rnbqkbnr/pppp1ppp/8/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R b KQkq - 1 2",
    },
    TestMoveStruct{
      "b8c6", "r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3",
    },
    TestMoveStruct{
      "f1b5", "r1bqkbnr/pppp1ppp/2n5/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R b KQkq - 3 3",
    },
    TestMoveStruct{
      "g8f6", "r1bqkb1r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 4 4",
    },
    TestMoveStruct{
      "e1g1", "r1bqkb1r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQ1RK1 b kq - 5 4",
    },
    TestMoveStruct{
      "d7d6", "r1bqkb1r/ppp2ppp/2np1n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQ1RK1 w kq - 0 5",
    },
    TestMoveStruct{
      "f1e1", "r1bqkb1r/ppp2ppp/2np1n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQR1K1 b kq - 1 5",
    },
    TestMoveStruct{
      "a7a6", "r1bqkb1r/1pp2ppp/p1np1n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQR1K1 w kq - 0 6",
    },
    TestMoveStruct{
      "b5a4", "r1bqkb1r/1pp2ppp/p1np1n2/4p3/B3P3/5N2/PPPP1PPP/RNBQR1K1 b kq - 1 6",
    },
    TestMoveStruct{
      "b7b5", "r1bqkb1r/2p2ppp/p1np1n2/1p2p3/B3P3/5N2/PPPP1PPP/RNBQR1K1 w kq b6 0 7",
    },
    TestMoveStruct{
      "a4b3", "r1bqkb1r/2p2ppp/p1np1n2/1p2p3/4P3/1B3N2/PPPP1PPP/RNBQR1K1 b kq - 1 7",
    },
    TestMoveStruct{
      "b5b4", "r1bqkb1r/2p2ppp/p1np1n2/4p3/1p2P3/1B3N2/PPPP1PPP/RNBQR1K1 w kq - 0 8",
    },
    TestMoveStruct{
      "c2c4", "r1bqkb1r/2p2ppp/p1np1n2/4p3/1pP1P3/1B3N2/PP1P1PPP/RNBQR1K1 b kq c3 0 8",
    },
    TestMoveStruct{
      "b4c3", "r1bqkb1r/2p2ppp/p1np1n2/4p3/4P3/1Bp2N2/PP1P1PPP/RNBQR1K1 w kq - 0 9",
    },
    TestMoveStruct{
      "b1c3", "r1bqkb1r/2p2ppp/p1np1n2/4p3/4P3/1BN2N2/PP1P1PPP/R1BQR1K1 b kq - 0 9",
    },
    TestMoveStruct{
      "f8e7", "r1bqk2r/2p1bppp/p1np1n2/4p3/4P3/1BN2N2/PP1P1PPP/R1BQR1K1 w kq - 1 10",
    },
    TestMoveStruct{
      "d2d4", "r1bqk2r/2p1bppp/p1np1n2/4p3/3PP3/1BN2N2/PP3PPP/R1BQR1K1 b kq d3 0 10",
    },
    TestMoveStruct{
      "e8g8", "r1bq1rk1/2p1bppp/p1np1n2/4p3/3PP3/1BN2N2/PP3PPP/R1BQR1K1 w - - 1 11",
    },
    TestMoveStruct{
      "d4e5", "r1bq1rk1/2p1bppp/p1np1n2/4P3/4P3/1BN2N2/PP3PPP/R1BQR1K1 b - - 0 11",
    },
    TestMoveStruct{
      "c6d4", "r1bq1rk1/2p1bppp/p2p1n2/4P3/3nP3/1BN2N2/PP3PPP/R1BQR1K1 w - - 1 12",
    },
    TestMoveStruct{
      "e5f6", "r1bq1rk1/2p1bppp/p2p1P2/8/3nP3/1BN2N2/PP3PPP/R1BQR1K1 b - - 0 12",
    },
    TestMoveStruct{
      "d4b3", "r1bq1rk1/2p1bppp/p2p1P2/8/4P3/1nN2N2/PP3PPP/R1BQR1K1 w - - 0 13",
    },
    TestMoveStruct{
      "f6e7", "r1bq1rk1/2p1Pppp/p2p4/8/4P3/1nN2N2/PP3PPP/R1BQR1K1 b - - 0 13",
    },
    TestMoveStruct{
      "b3c1", "r1bq1rk1/2p1Pppp/p2p4/8/4P3/2N2N2/PP3PPP/R1nQR1K1 w - - 0 14",
    },
    TestMoveStruct{
      "e7d8q", "r1bQ1rk1/2p2ppp/p2p4/8/4P3/2N2N2/PP3PPP/R1nQR1K1 b - - 0 14",
    },
    TestMoveStruct{
      "g8h8", "r1bQ1r1k/2p2ppp/p2p4/8/4P3/2N2N2/PP3PPP/R1nQR1K1 w - - 1 15",
    },
    TestMoveStruct{
      "d8f8", "r1b2Q1k/2p2ppp/p2p4/8/4P3/2N2N2/PP3PPP/R1nQR1K1 b - - 0 15",
    },
  }

  b := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  for _, test := range(tests) {
    b.moveUCI(test.uci)
    generatedFen := b.GenerateFen()
    if generatedFen != test.fen {
      t.Error(fmt.Sprintf("Expected fen %v after move %v, got %v", test.fen, test.uci, generatedFen))
    }
  }
}
