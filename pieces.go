package main

var SIZE int = 7

type Piece interface {
  GetPseudoLegalMoves(pos Position, b *Board) []Position
  GetLegalMoves(pos Position, b *Board) []Position
}

type DifferentIfMoved struct {
  hasMoved bool
}
