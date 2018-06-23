package main

import (
  "fmt"
  "strings"
  "time"
  "strconv"
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

var whiteSymbols = map[string]bool{
  "R": true,
  "N": true,
  "B": true,
  "Q": true,
  "K": true,
  "P": true,
}
var blackSymbols = map[string]bool{
  "r": true,
  "n": true,
  "b": true,
  "q": true,
  "k": true,
  "p": true,
}

func GenerateBoard(fen string) Board {
  b := Board{}
  fields := strings.Split(fen, " ")

  b.whites, b.blacks = GeneratePositions(fields[0])
  b.whiteToMove = (fields[1] == "w")
  b.availableCastles = SetAvailableCastles(fields[2])
  b.enPassantSquare = SetEnpassantSquare(fields[3])

  halfMoveClock, err := strconv.Atoi(fields[4])
  if err == nil {
    b.halfMoveClock = halfMoveClock
  }
  fullMoveNumber, err := strconv.Atoi(fields[5])
  if err == nil {
    b.fullMoveNumber = fullMoveNumber
  }

  return b
}

func SetAvailableCastles(avail string) map[string]bool {
  return map[string]bool {
     "bk" : strings.Contains(avail, "k"),
     "bq": strings.Contains(avail, "q"),
     "wk": strings.Contains(avail, "K"),
     "wq": strings.Contains(avail, "Q"),
  }
}

func SetEnpassantSquare(algebraic string) Position {
  row, col := -1, -1
  if len(algebraic) == 2 {
    cols := "abcdefgh"
    row = int(algebraic[1]) - 1
    col = strings.Index(cols, string(algebraic[0]))
  }
  return Position{row,col}
}

func GeneratePositions(pos string) (map[Position]string, map[Position]string) {
    rows := strings.Split(pos, "/")
    whites := make(map[Position]string)
    blacks := make(map[Position]string)

    for i, row := range rows {
      offset := 0
      for j, sq := range row {
        _, white := whiteSymbols[string(sq)]
        _, black := blackSymbols[string(sq)]
        switch {
        case white:
          whites[Position{7-i,j+offset}] = string(sq)
        case black:
          blacks[Position{7-i,j+offset}] = string(sq)
        default:
          // This gives the numeric value of the rune(ie '56' to the int 8)
          // Minus one because the offset number takes up a square itself
          offset += int(sq - '0') - 1
        }
      }
    }
    return whites, blacks
}



func main() {
  // fen := "rnbqk1nr/ppp2ppp/8/2bpp3/2B1P3/5N2/PPPP1PPP/RNBQ1RK1 b kq - 1 4"
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  start := time.Now()
  b := GenerateBoard(fen)
  // n := Knight{}
  // n.isWhite = true
  bish := Queen{}
  bish.isWhite = true
  res := bish.GetPseudoLegalMoves(Position{4,4}, &b)
  // // fmt.Println(n.isWhite)
  // res := n.GetPseudoLegalMoves(Position{0,0}, &b)
  fmt.Println(Position{4,4})
  fmt.Println(res)


  t := time.Now()
  fmt.Println(t.Sub(start))
}
