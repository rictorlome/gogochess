// This file covers bishop, rook, and queen
package main

type Slider struct {
  isWhite bool
  moveDiffs [][]int
}

func (s *Slider) slide(moveDiffs [][]int, pos Position, b *Board) []Position {
  var res []Position
  OUTER:
  for _, moveDiff := range(moveDiffs) {
    for i := 1; i <= SIZE; i++ {
      newPos := Position{pos.row + moveDiff[0] * i, pos.col + moveDiff[1] * i}
      // Cannot capture own piece
      if b.hasColoredPieceThere(s.isWhite, newPos) {
        continue OUTER
      }
      if newPos.isOnBoard() {
        res = append(res, newPos)
      }
      // Cannot slide beyond captured piece of opposite color
      if b.hasColoredPieceThere(!s.isWhite, newPos) {
        continue OUTER
      }
    }
  }
  return res
}

type Bishop struct {
  Slider
}

func (bish *Bishop) GetPseudoLegalMoves(pos Position, b *Board) []Position {
  moveDiffs := [][]int{
      []int{1,1}, []int{-1,-1},
      []int{-1,1}, []int{1,-1},
  }
  return bish.slide(moveDiffs, pos, b)
}

type Rook struct {
  Slider
  DifferentIfMoved
}

func (r *Rook) GetPseudoLegalMoves(pos Position, b *Board) []Position {
  moveDiffs := [][]int{
      []int{1,0}, []int{0,1},
      []int{-1,0}, []int{0,-1},
  }
  return r.slide(moveDiffs, pos, b)
}

type Queen struct {
  Slider
}

func (q *Queen) GetPseudoLegalMoves(pos Position, b *Board) []Position {
  moveDiffs := [][]int{
      []int{1,0}, []int{0,1},
      []int{-1,0}, []int{0,-1},
      []int{1,1}, []int{-1,-1},
      []int{-1,1}, []int{1,-1},
  }
  return q.slide(moveDiffs, pos, b)
}
