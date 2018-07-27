package main

var SIZE int = 7

var NULLMOVEDIFFS [][]int

type Piece interface {
	IsWhite() bool
	ToString() string
	CanPossiblyAttack(pos Position, target Position) (bool, [][]int)
	GetDefaultMoveDiffs() [][]int
	//AttackingSquares refer to squares which piece is currently attacking
	GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool
	//PseudoLegalMoves refer to unblocked moves on the board (including castle and en passant)
	// - this is the same as attacking squares for all pieces except pawn and king
	GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool
	//LegalMoves takes check into account
	GetLegalMoves(pos Position, b *Board) map[Position]bool
}
