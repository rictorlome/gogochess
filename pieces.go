package main

var SIZE int = 7

type Piece interface {
	IsWhite() bool
	ToString() string
	CanPossiblyAttack(pos Position, target Position) bool
	//AttackingSquares refer to squares which piece is currently attacking
	GetAttackingSquares(pos Position, b *Board) map[Position]bool
	//PseudoLegalMoves refer to unblocked moves on the board (including castle and en passant)
	// - this is the same as attacking squares for all pieces except pawn and king
	GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool
	//LegalMoves takes check into account
	GetLegalMoves(pos Position, b *Board) map[Position]bool
}
