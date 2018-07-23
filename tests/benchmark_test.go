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
    piece.GetAttackingSquares(pos, &board)
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
    piece.GetAttackingSquares(pos, &board)
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
  board := GenerateBoard("r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")
  for i := 0; i < b.N; i++ {
     board.GetAllLegalMoves(i % 2 == 0)
  }
}
