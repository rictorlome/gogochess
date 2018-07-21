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

func (k *King) GetAttackingSquares(pos Position, b *Board) map[Position]bool {
  res := make(map[Position]bool)
  nums := []int{-1,0,1}
  for _, i := range(nums) {
    for _, j := range(nums) {
      if !(i == 0 && j == 0) {
        newPos := Position{pos.row+i,pos.col+j}
        if newPos.isOnBoard() && !b.hasColoredPieceThere(k.isWhite, newPos) {
          res[newPos] = true
        }
      }
    }
  }
  return res
}
//Board state already knows whether king and rook have moved.
func (k *King) GetPseudoLegalMoves(pos Position, b *Board) map[Position]bool {
  result := k.GetAttackingSquares(pos, b)
  kingside, queenside := "bk", "bq"
  row := 7
  if k.isWhite {
    kingside, queenside = "wk", "wq"
    row = 0
  }
  queensquare, kingsquare := Position{row,2}, Position{row,6}
  interqueen := []Position{
    Position{row,1}, Position{row,2}, Position{row,3},
  }
  interking := []Position{
    Position{row,6}, Position{row,5},
  }
  if b.availableCastles[kingside] && b.areEmptySquares(interking) {
    result[kingsquare] = true
  }
  if b.availableCastles[queenside] && b.areEmptySquares(interqueen) {
    result[queensquare] = true
  }
  return result
}
