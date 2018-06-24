package main

type Knight struct {
  isWhite bool
}

func (n *Knight) GetMovesOnBoard(pos Position) []Position {
  var res []Position
  moveDiffs := [][]int{
    []int{1,2}, []int{1,-2},
    []int{-1,2}, []int{-1,-2},
    []int{2,1}, []int{2,-1},
    []int{-2,1}, []int{-2,-1},
  }
  for _, diff := range(moveDiffs) {
    newPos := Position{pos.row+diff[0],pos.col+diff[1]}
    if newPos.isOnBoard() {
      res = append(res, newPos)
    }
  }
  return res
}

func (n *Knight) GetPseudoLegalMoves(pos Position, b *Board) []Position {
  var res []Position
  movesOnBoard := n.GetMovesOnBoard(pos)
  ownColor := b.getColoredPieces(n.isWhite)
  for _, move := range(movesOnBoard) {
    _, pieceThere := ownColor[move]
    if !pieceThere {
      res = append(res, move)
    }
  }
  return res
}
