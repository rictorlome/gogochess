package main

type King struct {
  isWhite bool
  DifferentIfMoved
}

func (k *King) GetMovesOnBoard(pos Position) []Position {
  var res []Position
  nums := []int{-1,0,1}
  for _, i := range(nums) {
    for _, j := range(nums) {
      if !(i == 0 && j == 0) {
        newPos := Position{pos.row+i,pos.col+j}
        res = append(res, newPos)
      }
    }
  }
  return res
}
