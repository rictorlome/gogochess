package main

var knightMoveDiffs = [][]int{
	[]int{1, 2}, []int{1, -2},
	[]int{-1, 2}, []int{-1, -2},
	[]int{2, 1}, []int{2, -1},
	[]int{-2, 1}, []int{-2, -1},
}

type Knight struct {
	isWhite bool
}

func (n *Knight) ToString() string {
	if n.isWhite {
		return "N"
	}
	return "n"
}

func (n *Knight) IsWhite() bool {
	return n.isWhite
}

func (n *Knight) CanPossiblyAttack(pos Position, target Position) bool {
	rowDiff, colDiff := target.row-pos.row, target.col-pos.col
	return -2 <= rowDiff && rowDiff <= 2 && -2 <= colDiff && colDiff <= 2
}

func (n *Knight) GetDefaultMoveDiffs() [][]int {
	return knightMoveDiffs
}

func (n *Knight) GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool {
	res := make(map[Position]bool)

	for _, diff := range moveDiffs {
		newPos := Position{pos.row + diff[0], pos.col + diff[1]}
		if newPos.isOnBoard() {
			res[newPos] = true
		}
	}
	return res
}

func (n *Knight) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	result := n.GetAttackingSquares(pos, b, knightMoveDiffs)
	for move := range result {
		if b.hasColoredPieceThere(n.isWhite, move) {
			delete(result, move)
		}
	}
	return result
}

func (n *Knight) GetLegalMoves(pos Position, b *Board) map[Position]bool {
	result := n.GetPseudoLegalMoves(pos, b)
	for move := range result {
		if b.wouldCauseCheck(pos, move, "") {
			delete(result, move)
		}
	}
	return result
}
