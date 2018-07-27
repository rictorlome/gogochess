package main

import (
  "testing"
)

func BenchmarkParseFen(b *testing.B) {
  for i := 0; i < b.N; i++ {
    GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  }
}

func BenchmarkGenerateFen(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  for i := 0; i < b.N; i++ {
    board.GenerateFen()
  }
}
func BenchmarkGetColoredPieces(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  for i := 0; i < b.N; i++ {
    board.getColoredPieces(i % 2 == 0)
  }
}
func BenchmarkGetColoredKing(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  for i := 0; i < b.N; i++ {
    board.getColoredKing(i % 2 == 0)
  }
}
func BenchmarkHasColoredPieceThere(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  pos := Position{0,0}
  for i := 0; i < b.N; i++ {
    board.hasColoredPieceThere(i % 2 == 0, pos)
  }
}
func BenchmarkToPos(b *testing.B) {
  poses := []string{"a1","h8","-", "h7", "a8", "e4", "d4", "c2", "f8", "g3", "g1"}
  for i := 0; i < b.N; i++ {
    ToPos(poses[i % len(poses)])
  }
}
func BenchmarkToPiece(b *testing.B) {
  pieces := []string{"p","r","k","b","q","k","P","R","K","B","Q","K"}
  for i := 0; i < b.N; i++ {
    ToPiece(pieces[i % len(pieces)])
  }
}
func BenchmarkToStringPos(b *testing.B) {
  poses := []Position{
    Position{1,1}, Position{2,2}, Position{3,3}, Position{7,7},
    Position{-1,-1}, Position{6,0}, Position{4,6}, Position{2,7},
  }
  for i := 0; i < b.N; i++ {
    poses[i % len(poses)].String()
  }
}
func BenchmarkToStringPiece(b *testing.B) {
  pieces := []Piece{ToPiece("p"),ToPiece("r"),ToPiece("k"),ToPiece("b"),ToPiece("q"),ToPiece("k"),ToPiece("P"),ToPiece("R"),ToPiece("K"),ToPiece("B"),ToPiece("Q"),ToPiece("K")}
  for i := 0; i < b.N; i++ {
    pieces[i % len(pieces)].ToString()
  }
}
func BenchmarkFindPiece(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  poses := []Position{
     ToPos("a1"), ToPos("h8"), ToPos("-"), ToPos("h7"), ToPos("a8"),
     ToPos("e4"), ToPos("d4"), ToPos("c2"), ToPos("f8"), ToPos("g3"), ToPos("g1"),
   }
  for i := 0; i < b.N; i++ {
    board.findPiece(poses[i % len(poses)])
  }
}
func BenchmarkParseMove(b *testing.B) {
  for i := 0; i < b.N; i++ {
    parseMove("e2e4")
  }
}
func BenchmarkNaiveMove(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  start, end, promotion := parseMove("e2e4")
  for i := 0; i < b.N; i++ {
    if i % 2 == 0 {
      board.naiveMove(start,end,promotion)
    } else {
      board.naiveMove(end,start,promotion)
    }
  }
}

func BenchmarkGetAttackingSquaresInitial(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  var poses []Position
  for pos,_ := range(board.whites) {
    poses = append(poses,pos)
  }
  for pos,_ := range(board.blacks) {
    poses = append(poses,pos)
  }
  for i := 0; i < b.N; i++ {
    pos := poses[i % len(poses)]
    _, piece := board.findPiece(pos)
    piece.GetAttackingSquares(pos, &board, piece.GetDefaultMoveDiffs())
  }
}
func BenchmarkGetAttackingSquaresMiddle(b *testing.B) {
  board := GenerateBoard("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
  var poses []Position
  for pos,_ := range(board.whites) {
    poses = append(poses,pos)
  }
  for pos,_ := range(board.blacks) {
    poses = append(poses,pos)
  }
  for i := 0; i < b.N; i++ {
    pos := poses[i % len(poses)]
    _, piece := board.findPiece(pos)
    piece.GetAttackingSquares(pos, &board, piece.GetDefaultMoveDiffs())
  }
}
func BenchmarkGetPseudoLegalMovesInitial(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  var poses []Position
  for pos,_ := range(board.whites) {
    poses = append(poses,pos)
  }
  for pos,_ := range(board.blacks) {
    poses = append(poses,pos)
  }
  for i := 0; i < b.N; i++ {
    pos := poses[i % len(poses)]
    _, piece := board.findPiece(pos)
    piece.GetPseudoLegalMoves(pos, &board)
  }
}
func BenchmarkGetPseudoLegalMovesMiddle(b *testing.B) {
  board := GenerateBoard("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
  var poses []Position
  for pos,_ := range(board.whites) {
    poses = append(poses,pos)
  }
  for pos,_ := range(board.blacks) {
    poses = append(poses,pos)
  }
  for i := 0; i < b.N; i++ {
    pos := poses[i % len(poses)]
    _, piece := board.findPiece(pos)
    piece.GetPseudoLegalMoves(pos, &board)
  }
}
func BenchmarkGetLegalMovesInitial(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  var poses []Position
  for pos,_ := range(board.whites) {
    poses = append(poses,pos)
  }
  for pos,_ := range(board.blacks) {
    poses = append(poses,pos)
  }
  for i := 0; i < b.N; i++ {
    pos := poses[i % len(poses)]
    _, piece := board.findPiece(pos)
    piece.GetLegalMoves(pos, &board)
  }
}
func BenchmarkGetLegalMovesMiddle(b *testing.B) {
  board := GenerateBoard("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
  var poses []Position
  for pos,_ := range(board.whites) {
    poses = append(poses,pos)
  }
  for pos,_ := range(board.blacks) {
    poses = append(poses,pos)
  }
  for i := 0; i < b.N; i++ {
    pos := poses[i % len(poses)]
    _, piece := board.findPiece(pos)
    piece.GetLegalMoves(pos, &board)
  }
}

var result map[Position]bool

func BenchmarkGetAllLegalMovesInitial(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  for i := 0; i < b.N; i++ {
     board.GetAllLegalMoves(i % 2 == 0)
  }
}
func BenchmarkGetAllLegalMovesMiddle(b *testing.B) {
  boards := []Board {
    GenerateBoard("r3qb1k/1b4p1/p2pr2p/3n4/Pnp1N1N1/6RP/1B3PP1/1B1QR1K1 w - - bm Nxh6"),
    GenerateBoard("r4rk1/pp1n1p1p/1nqP2p1/2b1P1B1/4NQ2/1B3P2/PP2K2P/2R5 w - - bm Rxc5"),
    GenerateBoard("r2qk2r/ppp1b1pp/2n1p3/3pP1n1/3P2b1/2PB1NN1/PP4PP/R1BQK2R w KQkq - bm Nxg5"),
    GenerateBoard("r1b1kb1r/1p1n1ppp/p2ppn2/6BB/2qNP3/2N5/PPP2PPP/R2Q1RK1 w kq - bm Nxe6"),
    GenerateBoard("r2qrb1k/1p1b2p1/p2ppn1p/8/3NP3/1BN5/PPP3QP/1K3RR1 w - - bm e5"),
    GenerateBoard("rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kq - bm axb5"),
    GenerateBoard("1r1bk2r/2R2ppp/p3p3/1b2P2q/4QP2/4N3/1B4PP/3R2K1 w k - bm Rxd8+"),
    GenerateBoard("r3rbk1/ppq2ppp/2b1pB2/8/6Q1/1P1B3P/P1P2PP1/R2R2K1 w - - bm Bxh7+"),
    GenerateBoard("r4r1k/4bppb/2n1p2p/p1n1P3/1p1p1BNP/3P1NP1/qP2QPB1/2RR2K1 w - - bm Ng5"),
    GenerateBoard("r1b2rk1/1p1nbppp/pq1p4/3B4/P2NP3/2N1p3/1PP3PP/R2Q1R1K w - - bm Rxf7"),
    GenerateBoard("r1b3k1/p2p1nP1/2pqr1Rp/1p2p2P/2B1PnQ1/1P6/P1PP4/1K4R1 w - - bm Rxh6"),
  }
  for i := 0; i < b.N; i++ {
     boards[i % len(boards)].GetAllLegalMoves(i % 2 == 0)
  }
}

func BenchmarkInCheckmateSome(b *testing.B) {
  boards := []Board {
    GenerateBoard("r1b1k2r/ppp2ppp/8/4p3/2P5/3nnP2/PP1NN1PP/2R1K2R w kq - 0 15"),
    GenerateBoard("2q3k1/p1K2p1p/4qP1P/3p2P1/8/8/8/8 w - - 3 47"),
    GenerateBoard("2r1kb1R/1q1nQp2/p5p1/4p1B1/1p2P1P1/2N2P2/PPP5/2K5 b - - 1 23"),
    GenerateBoard("8/pp4pp/2p5/6Qk/1PP4P/P4pP1/5P1K/8 b - - 2 39"),
    GenerateBoard("8/6pk/3p3p/4p3/4Pn1P/8/6q1/7K w - - 0 55"),
    GenerateBoard("4R1k1/8/r4BK1/8/5P1p/6P1/7P/8 b - - 1 49"),
    GenerateBoard("8/8/7p/pp5P/3k4/Pb2p3/q7/K7 w - - 8 54"),
    GenerateBoard("1r2qNrk/2R3pQ/p6p/8/8/P6P/1P3PPK/8 b - - 8 36"),
    GenerateBoard("4r1k1/2p2p2/1rpb2pp/3p4/1P1Pp3/N1P5/P2N1P1q/R2Q1RK1 w - - 2 25"),
    // others' games
    GenerateBoard("R7/6pk/5p2/1P2bP1p/7K/3B2P1/7r/8 w - - 0 42"),
    GenerateBoard("rnbq3r/pp2kQpp/2pp4/2b1p1N1/4P3/3P4/PPP2PPP/RNB2RK1 b - - 3 10"),
  }
  for i := 0; i < b.N; i++ {
     boards[i % len(boards)].inCheckmate(i % 2 == 0)
  }
}

func BenchmarkInCheckmateNone(b *testing.B) {
  boards := []Board {
    GenerateBoard("r3qb1k/1b4p1/p2pr2p/3n4/Pnp1N1N1/6RP/1B3PP1/1B1QR1K1 w - - bm Nxh6"),
    GenerateBoard("r4rk1/pp1n1p1p/1nqP2p1/2b1P1B1/4NQ2/1B3P2/PP2K2P/2R5 w - - bm Rxc5"),
    GenerateBoard("r2qk2r/ppp1b1pp/2n1p3/3pP1n1/3P2b1/2PB1NN1/PP4PP/R1BQK2R w KQkq - bm Nxg5"),
    GenerateBoard("r1b1kb1r/1p1n1ppp/p2ppn2/6BB/2qNP3/2N5/PPP2PPP/R2Q1RK1 w kq - bm Nxe6"),
    GenerateBoard("r2qrb1k/1p1b2p1/p2ppn1p/8/3NP3/1BN5/PPP3QP/1K3RR1 w - - bm e5"),
    GenerateBoard("rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kq - bm axb5"),
    GenerateBoard("1r1bk2r/2R2ppp/p3p3/1b2P2q/4QP2/4N3/1B4PP/3R2K1 w k - bm Rxd8+"),
    GenerateBoard("r3rbk1/ppq2ppp/2b1pB2/8/6Q1/1P1B3P/P1P2PP1/R2R2K1 w - - bm Bxh7+"),
    GenerateBoard("r4r1k/4bppb/2n1p2p/p1n1P3/1p1p1BNP/3P1NP1/qP2QPB1/2RR2K1 w - - bm Ng5"),
    GenerateBoard("r1b2rk1/1p1nbppp/pq1p4/3B4/P2NP3/2N1p3/1PP3PP/R2Q1R1K w - - bm Rxf7"),
    GenerateBoard("r1b3k1/p2p1nP1/2pqr1Rp/1p2p2P/2B1PnQ1/1P6/P1PP4/1K4R1 w - - bm Rxh6"),
  }
  for i := 0; i < b.N; i++ {
     boards[i % len(boards)].inCheckmate(i % 2 == 0)
  }
}
func BenchmarkInCheckSome(b *testing.B) {
  boards := []Board {
    GenerateBoard("r1b1k2r/ppp2ppp/8/4p3/2P5/3nnP2/PP1NN1PP/2R1K2R w kq - 0 15"),
    GenerateBoard("2q3k1/p1K2p1p/4qP1P/3p2P1/8/8/8/8 w - - 3 47"),
    GenerateBoard("2r1kb1R/1q1nQp2/p5p1/4p1B1/1p2P1P1/2N2P2/PPP5/2K5 b - - 1 23"),
    GenerateBoard("8/pp4pp/2p5/6Qk/1PP4P/P4pP1/5P1K/8 b - - 2 39"),
    GenerateBoard("8/6pk/3p3p/4p3/4Pn1P/8/6q1/7K w - - 0 55"),
    GenerateBoard("4R1k1/8/r4BK1/8/5P1p/6P1/7P/8 b - - 1 49"),
    GenerateBoard("8/8/7p/pp5P/3k4/Pb2p3/q7/K7 w - - 8 54"),
    GenerateBoard("1r2qNrk/2R3pQ/p6p/8/8/P6P/1P3PPK/8 b - - 8 36"),
    GenerateBoard("4r1k1/2p2p2/1rpb2pp/3p4/1P1Pp3/N1P5/P2N1P1q/R2Q1RK1 w - - 2 25"),
    // others' games
    GenerateBoard("R7/6pk/5p2/1P2bP1p/7K/3B2P1/7r/8 w - - 0 42"),
    GenerateBoard("rnbq3r/pp2kQpp/2pp4/2b1p1N1/4P3/3P4/PPP2PPP/RNB2RK1 b - - 3 10"),
  }
  for i := 0; i < b.N; i++ {
     boards[i % len(boards)].inCheck(i % 2 == 0)
  }
}

func BenchmarkInCheckNone(b *testing.B) {
  boards := []Board {
    GenerateBoard("r3qb1k/1b4p1/p2pr2p/3n4/Pnp1N1N1/6RP/1B3PP1/1B1QR1K1 w - - bm Nxh6"),
    GenerateBoard("r4rk1/pp1n1p1p/1nqP2p1/2b1P1B1/4NQ2/1B3P2/PP2K2P/2R5 w - - bm Rxc5"),
    GenerateBoard("r2qk2r/ppp1b1pp/2n1p3/3pP1n1/3P2b1/2PB1NN1/PP4PP/R1BQK2R w KQkq - bm Nxg5"),
    GenerateBoard("r1b1kb1r/1p1n1ppp/p2ppn2/6BB/2qNP3/2N5/PPP2PPP/R2Q1RK1 w kq - bm Nxe6"),
    GenerateBoard("r2qrb1k/1p1b2p1/p2ppn1p/8/3NP3/1BN5/PPP3QP/1K3RR1 w - - bm e5"),
    GenerateBoard("rnbqk2r/1p3ppp/p7/1NpPp3/QPP1P1n1/P4N2/4KbPP/R1B2B1R b kq - bm axb5"),
    GenerateBoard("1r1bk2r/2R2ppp/p3p3/1b2P2q/4QP2/4N3/1B4PP/3R2K1 w k - bm Rxd8+"),
    GenerateBoard("r3rbk1/ppq2ppp/2b1pB2/8/6Q1/1P1B3P/P1P2PP1/R2R2K1 w - - bm Bxh7+"),
    GenerateBoard("r4r1k/4bppb/2n1p2p/p1n1P3/1p1p1BNP/3P1NP1/qP2QPB1/2RR2K1 w - - bm Ng5"),
    GenerateBoard("r1b2rk1/1p1nbppp/pq1p4/3B4/P2NP3/2N1p3/1PP3PP/R2Q1R1K w - - bm Rxf7"),
    GenerateBoard("r1b3k1/p2p1nP1/2pqr1Rp/1p2p2P/2B1PnQ1/1P6/P1PP4/1K4R1 w - - bm Rxh6"),
  }
  for i := 0; i < b.N; i++ {
     boards[i % len(boards)].inCheck(i % 2 == 0)
  }
}
