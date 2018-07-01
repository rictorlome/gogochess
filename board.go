package main

import (
  "fmt"
  "strings"
)

type Position struct {
  row, col int
}

func ToPos(s string) Position {
  cols := "abcdefgh"
  r := int(s[1] - '0') - 1
  return Position{r, strings.Index(cols,s[0:1])}
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

func ToPiece(s string) Piece {
  var symbolToPiece = map[string]Piece{
    "R": &Rook{true},
    "N": &Knight{true},
    "B": &Bishop{true},
    "Q": &Queen{true},
    "K": &King{true},
    "P": &Pawn{true},
    "r": &Rook{false},
    "n": &Knight{false},
    "b": &Bishop{false},
    "q": &Queen{false},
    "k": &King{false},
    "p": &Pawn{false},
  }
  return symbolToPiece[s]
}


func (b *Board) findPiece(p Position) (bool, Piece) {
  for k, v := range(b.whites) {
    if k == p {
      return true, ToPiece(v)
    }
  }
  for k, v := range(b.blacks) {
    if k == p {
      return true, ToPiece(v)
    }
  }
  return false, nil
}
