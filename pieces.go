package main

// import "fmt"

var SIZE int = 7

type Piece interface {
  GetPseudoLegalMoves(pos Position, b *Board) []Position
  GetLegalMoves(pos Position, b *Board) []Position
}

type King struct {
  isWhite bool
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
