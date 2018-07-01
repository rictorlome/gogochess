package main

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
  // kingside, queenside := "bk", "bq"
  // queensquare, kingsquare := Position{7,0}, Position{7,7}
  // if k.isWhite {
  //   kingside, queenside = "wk", "wq"
  //   queensquare, kingsquare = Position{0,0}, Position{0,7}
  // }
  //condition for castle is:
  //king has not moved
  //rook has not moved
  //intermediate squares are empty
  //king not in check, king not passing through check

  return minusCastles
}
