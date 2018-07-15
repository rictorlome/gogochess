package main

import (
  "fmt"
  "strings"
)

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
    return "-"
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
  whites, blacks map[Position]Piece
  whiteKing Position
  blackKing Position
  whiteToMove bool
  availableCastles map[string]bool
  enPassantSquare Position
  halfMoveClock int
  fullMoveNumber int
}

func (b *Board) getColoredPieces(white bool) map[Position]Piece {
  if white {
    return b.whites
  }
  return b.blacks
}

func (b *Board) getColoredKing(white bool) Position {
  if white {
    return b.whiteKing
  }
  return b.blackKing
}

func (b *Board) hasColoredPieceThere(white bool, sq Position) bool {
  pieces := b.getColoredPieces(white)
  _, there := pieces[sq]
  return there
}

func ToPiece(s string) Piece {
  return symbolToPiece[s]
}

func (b *Board) inCheck(white bool) bool {
  kingPos := b.getColoredKing(white)
  attackingPieces := b.getColoredPieces(!white)
  for pos, piece := range(attackingPieces) {
    attackingSquares := piece.GetAttackingSquares(pos, b)
    for _, sq := range(attackingSquares) {
      if kingPos == sq {
        return true
      }
    }
  }
  return false
}


func (b *Board) findPiece(p Position) (bool, Piece) {
  piece := b.whites[p]
  if piece != nil {
    return true, piece
  }
  piece = b.blacks[p]
  if piece != nil {
    return true, piece
  }
  return false, nil
}
// Accepts UCI move format.
func parseMove(s string) (start Position, end Position, promotion string) {
  start = ToPos(s[0:2])
  end = ToPos(s[2:4])
  promotion = ""
  if len(s) == 5 {
    promotion = string(s[4])
  }
  return
}

func (b *Board) naiveMove(start Position, end Position, promotion string) {
  _, piece := b.findPiece(start)
  if piece.IsWhite() {
    delete(b.whites, start)
    b.whites[end] = piece
  } else {
    delete(b.blacks, start)
    b.blacks[end] = piece
  }
}

func (b *Board) wouldCauseCheck(start Position, end Position, promotion string) bool {
  _, piece := b.findPiece(start)
  occupied, target := b.findPiece(end)
  b.naiveMove(start,end,promotion)
  check := b.inCheck(piece.IsWhite())
  b.naiveMove(end,start,promotion)
  if occupied {
    pieces := b.getColoredPieces(!piece.IsWhite())
    pieces[end] = target
  }
  return check
}
