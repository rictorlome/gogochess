package main

import (
  "fmt"
)

type Position struct {
  row, col int
}

func (p Position) String() string {
  if p.col == -1 {
    return "nullPos"
  }
  cols := "abcdefgh"
  return fmt.Sprintf("%c%d", cols[p.col], p.row+1)
}

func (p Position) isOnBoard() bool {
  return 0 <= p.col && p.col <= 7 && 0 <= p.row && p.row <= 7
}

func (p Position) isEmpty(b *Board) bool {
  _, whiteThere := b.whites[p]
  _, blackThere := b.blacks[p]
  return (!whiteThere && !blackThere)
}

func (p Position) getNeighbors() []Position {
  return []Position {
    Position{p.row,p.col-1},
    Position{p.row,p.col+1},
  }
}

type Board struct {
  whites, blacks map[Position]string
  whiteToMove bool
  availableCastles map[string]bool
  enPassantSquare Position
  halfMoveClock int
  fullMoveNumber int
}

func (b *Board) getColoredPieces(white bool) map[Position]string {
  if white {
    return b.whites
  }
  return b.blacks
}

func (b *Board) hasColoredPieceThere(white bool, sq Position) bool {
  pieces := b.getColoredPieces(white)
  _, there := pieces[sq]
  return there
}
