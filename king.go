package main

var kingMoveDiffs = [][]int{
	[]int{1, -1}, []int{1, 0}, []int{1, 1},
	[]int{0, -1}, []int{0, 1},
	[]int{-1, -1}, []int{-1, 0}, []int{-1, 1},
}

type King struct {
	isWhite bool
}

func (k *King) IsWhite() bool {
	return k.isWhite
}

func (k *King) ToString() string {
	if k.isWhite {
		return "K"
	}
	return "k"
}

func (k *King) HomeRow() int {
	if k.isWhite {
		return 0
	}
	return 7
}

func (k *King) CanPossiblyAttack(pos Position, target Position) bool {
	rowDiff, colDiff := target.row-pos.row, target.col-pos.col
	return -1 <= rowDiff && rowDiff <= 1 && -1 <= colDiff && colDiff <= 1
}

func (k *King) GetDefaultMoveDiffs() [][]int {
	return kingMoveDiffs
}

func (k *King) GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool {
	res := make(map[Position]bool)
	for _, moveDiff := range moveDiffs {
		newPos := Position{pos.row + moveDiff[0], pos.col + moveDiff[1]}
		if newPos.isOnBoard() {
			res[newPos] = true
		}
	}
	return res
}

//Board state already knows whether king and rook have moved.
func (k *King) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	result := k.GetAttackingSquares(pos, b, kingMoveDiffs)
	for move, _ := range result {
		if b.hasColoredPieceThere(k.isWhite, move) {
			delete(result, move)
		}
	}
	row := k.HomeRow()
	kingside, queenside := "bk", "bq"
	if k.isWhite {
		kingside, queenside = "wk", "wq"
	}
	queensquare, kingsquare := Position{row, 2}, Position{row, 6}
	interqueen := []Position{
		Position{row, 1}, Position{row, 2}, Position{row, 3},
	}
	interking := []Position{
		Position{row, 6}, Position{row, 5},
	}
	if b.availableCastles[kingside] && b.areEmptySquares(interking) {
		result[kingsquare] = true
	}
	if b.availableCastles[queenside] && b.areEmptySquares(interqueen) {
		result[queensquare] = true
	}
	return result
}

func (k *King) GetLegalMoves(pos Position, b *Board) map[Position]bool {
	result := k.GetPseudoLegalMoves(pos, b)
	for move := range result {
		if b.wouldCauseCheck(pos, move, "") {
			delete(result, move)
		}
	}
	row := k.HomeRow()
	queensquare, kingsquare := Position{row, 2}, Position{row, 6}
	interqueen := []Position{
		Position{row, 2}, Position{row, 3}, Position{row, 4},
	}
	interking := []Position{
		Position{row, 6}, Position{row, 5}, Position{row, 4},
	}
	if result[queensquare] && b.areAttackedByColor(!k.isWhite, interqueen) {
		delete(result, queensquare)
	}
	if result[kingsquare] && b.areAttackedByColor(!k.isWhite, interking) {
		delete(result, kingsquare)
	}
	return result
}
