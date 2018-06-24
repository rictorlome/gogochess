package main

type King struct {
  isWhite bool
  DifferentIfMoved
}

func (k *King) GetAttackingSquares(pos Position, b *Board) []Position {
  var res []Position
  nums := []int{-1,0,1}
  for _, i := range(nums) {
    for _, j := range(nums) {
      if !(i == 0 && j == 0) {
        newPos := Position{pos.row+i,pos.col+j}
        if newPos.isOnBoard() && !b.hasColoredPieceThere(k.isWhite, newPos) {
          res = append(res, newPos)
        }
      }
    }
  }
  return res
}
//Temporary...
func (k *King) GetPseudoLegalMoves(pos Position, b *Board) []Position {
  minusCastles := k.GetAttackingSquares(pos, b)
  return minusCastles
}
