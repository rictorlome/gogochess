// This file covers bishop, rook, and queen
package main

var bishopMoveDiffs = [][]int{
	[]int{1, 1}, []int{-1, -1},
	[]int{-1, 1}, []int{1, -1},
}

var rookMoveDiffs = [][]int{
	[]int{1, 0}, []int{0, 1},
	[]int{-1, 0}, []int{0, -1},
}

var queenMoveDiffs = [][]int{
	[]int{1, 0}, []int{0, 1},
	[]int{-1, 0}, []int{0, -1},
	[]int{1, 1}, []int{-1, -1},
	[]int{-1, 1}, []int{1, -1},
}

func slide(isWhite bool, moveDiffs [][]int, pos Position, b *Board) map[Position]bool {
	res := make(map[Position]bool)
OUTER:
	for _, moveDiff := range moveDiffs {
		for i := 1; i <= SIZE; i++ {
			newPos := Position{pos.row + moveDiff[0]*i, pos.col + moveDiff[1]*i}
			if newPos.isOnBoard() {
				res[newPos] = true
			} else {
				continue OUTER
			}
			there, _ := b.findPiece(newPos)
			if there {
				continue OUTER
			}
		}
	}
	return res
}

type Bishop struct {
	isWhite bool
}

func (b *Bishop) ToString() string {
	if b.isWhite {
		return "B"
	}
	return "b"
}

func (b *Bishop) IsWhite() bool {
	return b.isWhite
}

func (bish *Bishop) CanPossiblyAttack(pos Position, target Position) (bool, [][]int) {
	// on top-left,bottom-right diagonal
	if pos.row+pos.col == target.row+target.col {
		if pos.row < target.row {
			return true, [][]int{[]int{1, -1}}
		}
		return true, [][]int{[]int{-1, 1}}
	}
	// on bottom-left, top-right diagonal
	if pos.row-pos.col == target.row-target.col {
		if pos.row < target.row {
			return true, [][]int{[]int{1, 1}}
		}
		return true, [][]int{[]int{-1, -1}}
	}
	return false, NULLMOVEDIFFS
}

func (bish *Bishop) GetDefaultMoveDiffs() [][]int {
	return bishopMoveDiffs
}

func (bish *Bishop) GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool {
	return slide(bish.isWhite, moveDiffs, pos, b)
}

func (bish *Bishop) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	result := bish.GetAttackingSquares(pos, b, bishopMoveDiffs)
	for move := range result {
		if b.hasColoredPieceThere(bish.isWhite, move) {
			delete(result, move)
		}
	}
	return result
}

func (bish *Bishop) GetLegalMoves(pos Position, b *Board) map[Position]bool {
	result := bish.GetPseudoLegalMoves(pos, b)
	for move := range result {
		if b.wouldCauseCheck(pos, move, "") {
			delete(result, move)
		}
	}
	return result
}

type Rook struct {
	isWhite bool
}

func (r *Rook) ToString() string {
	if r.isWhite {
		return "R"
	}
	return "r"
}

func (r *Rook) IsWhite() bool {
	return r.isWhite
}

func (r *Rook) CanPossiblyAttack(pos Position, target Position) (bool, [][]int) {
	if pos.row == target.row {
		if pos.col < target.col {
			return true, [][]int{[]int{0, 1}}
		}
		return true, [][]int{[]int{0, -1}}
	}
	if pos.col == target.col {
		if pos.row < target.row {
			return true, [][]int{[]int{1, 0}}
		}
		return true, [][]int{[]int{-1, 0}}
	}
	return false, NULLMOVEDIFFS
}

func (r *Rook) GetDefaultMoveDiffs() [][]int {
	return rookMoveDiffs
}

func (r *Rook) GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool {
	return slide(r.isWhite, moveDiffs, pos, b)
}

func (r *Rook) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	result := r.GetAttackingSquares(pos, b, rookMoveDiffs)
	for move := range result {
		if b.hasColoredPieceThere(r.isWhite, move) {
			delete(result, move)
		}
	}
	return result
}

func (r *Rook) GetLegalMoves(pos Position, b *Board) map[Position]bool {
	result := r.GetPseudoLegalMoves(pos, b)
	for move := range result {
		if b.wouldCauseCheck(pos, move, "") {
			delete(result, move)
		}
	}
	return result
}

type Queen struct {
	isWhite bool
}

func (q *Queen) ToString() string {
	if q.isWhite {
		return "Q"
	}
	return "q"
}

func (q *Queen) IsWhite() bool {
	return q.isWhite
}

func (q *Queen) CanPossiblyAttack(pos Position, target Position) (bool, [][]int) {
	if pos.row+pos.col == target.row+target.col {
		if pos.row < target.row {
			return true, [][]int{[]int{1, -1}}
		}
		return true, [][]int{[]int{-1, 1}}
	}
	// on bottom-left, top-right diagonal
	if pos.row-pos.col == target.row-target.col {
		if pos.row < target.row {
			return true, [][]int{[]int{1, 1}}
		}
		return true, [][]int{[]int{-1, -1}}
	}
	if pos.row == target.row {
		if pos.col < target.col {
			return true, [][]int{[]int{0, 1}}
		}
		return true, [][]int{[]int{0, -1}}
	}
	if pos.col == target.col {
		if pos.row < target.row {
			return true, [][]int{[]int{1, 0}}
		}
		return true, [][]int{[]int{-1, 0}}
	}
	return false, NULLMOVEDIFFS
}

func (q *Queen) GetDefaultMoveDiffs() [][]int {
	return queenMoveDiffs
}

func (q *Queen) GetAttackingSquares(pos Position, b *Board, moveDiffs [][]int) map[Position]bool {
	return slide(q.isWhite, moveDiffs, pos, b)
}

func (q *Queen) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
	result := q.GetAttackingSquares(pos, b, queenMoveDiffs)
	for move := range result {
		if b.hasColoredPieceThere(q.isWhite, move) {
			delete(result, move)
		}
	}
	return result
}

func (q *Queen) GetLegalMoves(pos Position, b *Board) map[Position]bool {
	result := q.GetPseudoLegalMoves(pos, b)
	for move := range result {
		if b.wouldCauseCheck(pos, move, "") {
			delete(result, move)
		}
	}
	return result
}
