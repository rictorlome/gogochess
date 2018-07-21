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
  if s == "-" {
    return Position{-1,-1}
  }
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
  capture, _ := b.findPiece(end)
  if piece.IsWhite() {
    delete(b.whites, start)
    b.whites[end] = piece
    if capture {
      delete(b.blacks, end)
    }
  } else {
    delete(b.blacks, start)
    b.blacks[end] = piece
    if capture {
      delete(b.whites, end)
    }
  }
}
// must be called before the actual move, in order to check if capture took place.
func (b *Board) updateBoardState(start Position, end Position) {
  _, piece := b.findPiece(start)
  // king position and castle
  if piece.ToString() == "K" {
    b.whiteKing = end
    b.availableCastles["wk"], b.availableCastles["wq"] = false, false
  }
  if piece.ToString() == "k" {
    b.blackKing = end
    b.availableCastles["bk"], b.availableCastles["bq"] = false, false
  }
  // en passant square (per lichess this only updates when opposite color has attacking pawn posted correctly)
  if piece.ToString() == "P" && start.row == 1 && end.row == 3 {
    b.enPassantSquare = Position{end.row-1,end.col}
  } else if piece.ToString() == "p" && start.row == 6 && end.row == 4 {
    b.enPassantSquare = Position{end.row+1,end.col}
  } else {
    b.enPassantSquare = Position{-1,-1}
  }
  // who to move
  b.whiteToMove = !b.whiteToMove
  // availabe castles
  if piece.ToString() == "R" {
    if start.col == 0 {
      b.availableCastles["wq"] = false
    }
    if start.col == 7 {
      b.availableCastles["wk"] = false
    }
  }
  if piece.ToString() == "r" {
    if start.col == 0 {
      b.availableCastles["bq"] = false
    }
    if start.col == 7 {
      b.availableCastles["bk"] = false
    }
  }
  // half clock move (number of moves since last pawn advance or capture)
  capture, _ := b.findPiece(end)
  if piece.ToString() == "P" || piece.ToString() == "p" || capture {
    b.halfMoveClock = 0
  } else {
    b.halfMoveClock += 1
  }
  // full move number (updated after whiteToMove is updated)
  if b.whiteToMove {
    b.fullMoveNumber += 1
  }
}
// must be called before update board state, in order to access previous en passant square
func (b *Board) cleanUpEnPassant(start Position, end Position) {
  there, piece := b.findPiece(start)
  if !there {
    return
  }
  if piece.ToString() == "P" && end == b.enPassantSquare {
    delete(b.blacks, Position{start.row,end.col})
  }
  if piece.ToString() == "p" && end == b.enPassantSquare {
    delete(b.whites, Position{start.row,end.col})
  }
}

func (b *Board) cleanUpCastle(start Position, end Position) {
  there, piece := b.findPiece(start)
  if !there {
    return
  }
  if piece.ToString() == "K" {
    //kingside
    if end.col > start.col + 1 {
      b.naiveMove(parseMove("h1f1"))
    }
    //queenside
    if end.col < start.col - 1 {
      b.naiveMove(parseMove("a1d1"))
    }
  }
  if piece.ToString() == "k" {
    //kingside
    if end.col > start.col + 1 {
      b.naiveMove(parseMove("h8f8"))
    }
    //queenside
    if end.col < start.col - 1 {
      b.naiveMove(parseMove("a8d8"))
    }
  }
}

func (b *Board) cleanUpPromotion(start Position, promotion string) {
  if promotion == "" {
    return
  }
  _, piece := b.findPiece(start)
  if piece.IsWhite() {
    b.whites[start] = ToPiece(strings.ToUpper(promotion))
  } else {
    b.blacks[start] = ToPiece(strings.ToLower(promotion))
  }
}

func (b *Board) move(s string) {
  start, end, promotion := parseMove(s)
  b.cleanUpEnPassant(start, end)
  b.cleanUpCastle(start, end)
  b.updateBoardState(start, end)
  b.cleanUpPromotion(start, promotion)
  b.naiveMove(start, end, promotion)
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
