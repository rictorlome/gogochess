package main

import (
  "strings"
  "strconv"
)

func GenerateBoard(fen string) Board {
  b := Board{}
  fields := strings.Split(fen, " ")

  b.whites, b.blacks, b.whiteKing, b.blackKing = GeneratePositions(fields[0])
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
     "bk": strings.Contains(avail, "k"),
     "bq": strings.Contains(avail, "q"),
     "wk": strings.Contains(avail, "K"),
     "wq": strings.Contains(avail, "Q"),
  }
}

func SetEnpassantSquare(algebraic string) Position {
  row, col := -1, -1
  if len(algebraic) == 2 {
    cols := "abcdefgh"
    row = int(algebraic[1] - '0') - 1
    col = strings.Index(cols, string(algebraic[0]))
  }
  return Position{row,col}
}

func GeneratePositions(pos string) (map[Position]Piece, map[Position]Piece, Position, Position) {
    rows := strings.Split(pos, "/")
    whites := make(map[Position]Piece)
    blacks := make(map[Position]Piece)
    var whiteKing, blackKing Position
    for i, row := range rows {
      offset := 0
      for j, sq := range row {
        piece := ToPiece(string(sq))
        if piece == nil {
          // This gives the numeric value of the rune(ie '56' to the int 8)
          // Minus one because the offset number takes up a square itself
          offset += int(sq - '0') - 1
        } else {
          pos := Position{7-i,j+offset}
          if piece.IsWhite() {
            whites[pos] = piece
            if piece.ToString() == "K" {
              whiteKing = pos
            }
          } else  {
            blacks[pos] = piece
            if piece.ToString() == "k" {
              blackKing = pos
            }
          }
        }
      }
    }
    return whites, blacks, whiteKing, blackKing
}

func (b *Board) GenerateFen() string {
  var fenArr []string
  //position string
  fenArr = append(fenArr, b.GenerateFenPositionString())
  if b.whiteToMove {
    fenArr = append(fenArr, "w")
  } else {
    fenArr = append(fenArr, "b")
  }
  fenArr = append(fenArr, b.GenerateCastleString())
  fenArr = append(fenArr, b.enPassantSquare.String())
  fenArr = append(fenArr, strconv.Itoa(b.halfMoveClock))
  fenArr = append(fenArr, strconv.Itoa(b.fullMoveNumber))
  return strings.Join(fenArr, " ")
}

func (b *Board) GenerateCastleString() string {
  var castleString string
  if b.availableCastles["wk"] { castleString += "K"}
  if b.availableCastles["wq"] { castleString += "Q"}
  if b.availableCastles["bk"] { castleString += "k"}
  if b.availableCastles["bq"] { castleString += "q"}
  if castleString == "" {castleString += "-"}
  return castleString
}

func (b *Board) GenerateFenPositionString() string{
  var grid [8][8]string
  var fenArr []string
  for pos, piece := range b.whites {
    grid[7-pos.row][pos.col] = piece.ToString()
  }
  for pos, piece := range b.blacks {
    grid[7-pos.row][pos.col] = piece.ToString()
  }
  for _, row := range grid {
    fenString := ""
    offset := 0
    for _, sq := range row {
      if sq == "" {
        offset += 1
      } else {
        if offset == 0 {
          fenString += sq
        } else {
          fenString += strconv.Itoa(offset) + sq
          offset = 0
        }
      }
    }
    if offset != 0 {
      fenString += strconv.Itoa(offset)
    }
    fenArr = append(fenArr,fenString)
  }
  return strings.Join(fenArr, "/")
}
