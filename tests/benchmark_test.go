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
func BenchmarkToPiece(b *testing.B) {
  pieces := []string{"p","r","k","b","q","k","P","R","K","B","Q","K"}
  for i := 0; i < b.N; i++ {
    ToPiece(pieces[i % len(pieces)])
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

var result map[Position]bool

func BenchmarkGetAllLegalMoves(b *testing.B) {
  board := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
  for i := 0; i < b.N; i++ {
     board.GetAllLegalMoves(i % 2 == 0)
  }
}
