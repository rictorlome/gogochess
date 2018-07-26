package main

var pawnMoveDiffs = [][]int{
	[]int{0, 1}, []int{0, -1},
}

type Pawn struct {
	isWhite bool
}

func (p *Pawn) IsWhite() bool {
	return p.isWhite
}

func (p *Pawn) ToString() string {
	if p.isWhite {
		return "P"
	}
	return "p"
}

func (p *Pawn) ForwardDir() int {
	if p.isWhite {
		return 1
	}
	return -1
}

func (p *Pawn) hasMoved(pos Position) bool {
	return (p.isWhite && pos.row != 1) || (!p.isWhite && pos.row != 6)
}

func (p *Pawn) CanPossiblyAttack(pos Position, target Position) bool {
	colDiff := target.col - pos.col
	return pos.row+p.ForwardDir() == target.row && (colDiff == 1 || colDiff == -1)
}

func (p *Pawn) GetDefaultMoveDiffs() [][]int {
	return pawnMoveDiffs
}

func (p *Pawn) GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool {
	res := make(map[Position]bool)
	for _, moveDiff := range moveDiffs {
		side := Position{pos.row + p.ForwardDir(), pos.col + moveDiff[1]}
		if side.isOnBoard() {
			res[side] = true
		}
	}
	return res
}

func (p *Pawn) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	res := p.GetAttackingSquares(pos, b, pawnMoveDiffs)
	for move := range res {
		if !b.hasColoredPieceThere(!p.isWhite, move) && b.enPassantSquare != move {
			delete(res, move)
		}
	}
	//can advance one into empty square
	forwardSquare := Position{pos.row + p.ForwardDir(), pos.col}
	if forwardSquare.isOnBoard() && forwardSquare.isEmpty(b) {
		res[forwardSquare] = true
	}
	//if has not moved, can advance two through empty into empty
	twoUp := Position{forwardSquare.row + p.ForwardDir(), forwardSquare.col}
	if !p.hasMoved(pos) && forwardSquare.isEmpty(b) && twoUp.isEmpty(b) {
		res[twoUp] = true
	}
	return res
}

func (p *Pawn) GetLegalMoves(pos Position, b *Board) map[Position]bool {
	result := p.GetPseudoLegalMoves(pos, b)
	for move := range result {
		if b.wouldCauseCheck(pos, move, "") {
			delete(result, move)
		}
	}
	return result
}
