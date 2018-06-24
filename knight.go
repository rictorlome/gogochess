package main

type Knight struct {
  isWhite bool
}

func (n *Knight) GetAttackingSquares(pos Position, b *Board) []Position {
  var res []Position
  moveDiffs := [][]int{
    []int{1,2}, []int{1,-2},
    []int{-1,2}, []int{-1,-2},
    []int{2,1}, []int{2,-1},
    []int{-2,1}, []int{-2,-1},
  }
  for _, diff := range(moveDiffs) {
    newPos := Position{pos.row+diff[0],pos.col+diff[1]}
    if newPos.isOnBoard() && !b.hasColoredPieceThere(n.isWhite, newPos) {
      res = append(res, newPos)
    }
  }
  return res
}

func (n *Knight) GetPseudoLegalMoves(pos Position, b *Board) []Position {
  return n.GetAttackingSquares(pos, b)
}
