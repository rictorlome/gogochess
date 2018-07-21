package main

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

func (n *Knight) GetAttackingSquares(pos Position, b *Board) map[Position]bool {
  res := make(map[Position]bool)
  moveDiffs := [][]int{
    []int{1,2}, []int{1,-2},
    []int{-1,2}, []int{-1,-2},
    []int{2,1}, []int{2,-1},
    []int{-2,1}, []int{-2,-1},
  }
  for _, diff := range(moveDiffs) {
    newPos := Position{pos.row+diff[0],pos.col+diff[1]}
    if newPos.isOnBoard() && !b.hasColoredPieceThere(n.isWhite, newPos) {
      res[newPos] = true
    }
  }
  return res
}

func (n *Knight) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
  return n.GetAttackingSquares(pos, b)
}
