// This file covers bishop, rook, and queen
package main


func slide(isWhite bool, moveDiffs [][]int, pos Position, b *Board) map[Position]bool {
  res := make(map[Position]bool)
  OUTER:
  for _, moveDiff := range(moveDiffs) {
    for i := 1; i <= SIZE; i++ {
      newPos := Position{pos.row + moveDiff[0] * i, pos.col + moveDiff[1] * i}
      // Cannot capture own piece
      if b.hasColoredPieceThere(isWhite, newPos) {
        continue OUTER
      }
      if newPos.isOnBoard() {
        res[newPos] = true
      }
      // Cannot slide beyond captured piece of opposite color
      if b.hasColoredPieceThere(!isWhite, newPos) {
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

func (bish *Bishop) GetAttackingSquares(pos Position, b *Board) map[Position]bool {
  moveDiffs := [][]int{
      []int{1,1}, []int{-1,-1},
      []int{-1,1}, []int{1,-1},
  }
  return slide(bish.isWhite, moveDiffs, pos, b)
}

func (bish *Bishop) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
  return bish.GetAttackingSquares(pos, b)
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

func (r *Rook) GetAttackingSquares(pos Position, b *Board) map[Position]bool {
  moveDiffs := [][]int{
      []int{1,0}, []int{0,1},
      []int{-1,0}, []int{0,-1},
  }
  return slide(r.isWhite, moveDiffs, pos, b)
}


func (r *Rook) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
  return r.GetAttackingSquares(pos, b)
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

func (q *Queen) GetAttackingSquares(pos Position, b *Board) map[Position]bool {
  moveDiffs := [][]int{
      []int{1,0}, []int{0,1},
      []int{-1,0}, []int{0,-1},
      []int{1,1}, []int{-1,-1},
      []int{-1,1}, []int{1,-1},
  }
  return slide(q.isWhite, moveDiffs, pos, b)
}

func (q *Queen) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
  return q.GetAttackingSquares(pos, b)
}
