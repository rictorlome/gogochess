package main

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

func (p *Pawn) hasMoved(pos Position) bool {
	return (p.isWhite && pos.row != 1) || (!p.isWhite && pos.row != 6)
}

func (p *Pawn) GetAttackingSquares(pos Position, b *Board) map[Position]bool {
	res := make(map[Position]bool)
	forwardDir := map[bool]int{
		true:  1,
		false: -1,
	}
	forwardSquare := Position{pos.row + forwardDir[p.isWhite], pos.col}
	//can capture opposite color (or empassant sq) if neighbors of above square
	for _, side := range forwardSquare.getNeighbors() {
		if side.isOnBoard() {
			res[side] = true
		}
	}
	return res
}

func (p *Pawn) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	res := p.GetAttackingSquares(pos, b)
	for move := range res {
		if !b.hasColoredPieceThere(!p.isWhite, move) && b.enPassantSquare != move {
			delete(res, move)
		}
	}
	forwardDir := map[bool]int{
		true:  1,
		false: -1,
	}
	//can advance one into empty square
	forwardSquare := Position{pos.row + forwardDir[p.isWhite], pos.col}
	if forwardSquare.isOnBoard() && forwardSquare.isEmpty(b) {
		res[forwardSquare] = true
	}
	//if has not moved, can advance two through empty into empty
	twoUp := Position{forwardSquare.row + forwardDir[p.isWhite], forwardSquare.col}
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
