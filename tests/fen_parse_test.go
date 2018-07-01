package main

import (
  "testing"
  "strings"
  "fmt"
)

type testpair struct {
  jsonPosition string
  FEN string
}

var fenPieceToObj = map[string]string {
  "R": "wR",
  "N": "wN",
  "B": "wB",
  "Q": "wQ",
  "K": "wK",
  "P": "wP",
  "r": "bR",
  "n": "bN",
  "b": "bB",
  "q": "bQ",
  "k": "bK",
  "p": "bP",
}

func jsonToMap(s string) map[string]string {
  posToPieces := strings.Split(s, ",")
  resMap := make(map[string]string)
  for _, posToPiece := range posToPieces {
    posAndPiece := strings.Split(posToPiece, ":")
    pos, piece := posAndPiece[0], posAndPiece[1]
    resMap[pos] = piece
  }
  return resMap
}

// These test cases were produced using the chessboard.js API
var jsonStringOne string = "a8:bR,b8:bN,c8:bB,d8:bQ,e8:bK,f8:bB,g8:bN,h8:bR,a7:bP,b7:bP,c7:bP,d7:bP,e7:bP,f7:bP,g7:bP,h7:bP,a2:wP,b2:wP,c2:wP,d2:wP,e2:wP,f2:wP,g2:wP,h2:wP,a1:wR,b1:wN,c1:wB,d1:wQ,e1:wK,f1:wB,g1:wN,h1:wR"
var fenOne string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var jsonStringTwo string = "a8:bR,b8:bN,c8:bB,d8:bQ,e8:bK,f8:bB,g8:bN,h8:bR,a7:bP,b7:bP,c7:bP,d7:bP,f7:bP,g7:bP,h7:bP,a2:wP,b2:wP,c2:wP,d2:wP,f2:wP,g2:wP,h2:wP,a1:wR,b1:wN,c1:wB,d1:wQ,e1:wK,f1:wB,g1:wN,h1:wR,e4:wP,e5:bP"
var fenTwo string = "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2"

var jsonStringThree string = "a8:bR,c8:bB,d8:bQ,e8:bK,f8:bB,g8:bN,h8:bR,a7:bP,b7:bP,c7:bP,d7:bP,f7:bP,g7:bP,h7:bP,a2:wP,b2:wP,c2:wP,f2:wP,g2:wP,h2:wP,a1:wR,b1:wN,c1:wB,d1:wQ,e1:wK,f1:wB,h1:wR,e4:wP,f3:wN,c6:bN,d4:bP"
var fenThree string = "r1bqkbnr/pppp1ppp/2n5/8/3pP3/5N2/PPP2PPP/RNBQKB1R w KQkq - 0 4"

var jsonStringFour string = "a8:bR,b8:bN,c8:bB,d8:bQ,e8:bK,g8:bN,h8:bR,a7:bP,b7:bP,c7:bP,f7:bP,g7:bP,h7:bP,a2:wP,b2:wP,c2:wP,d2:wP,f2:wP,g2:wP,h2:wP,a1:wR,b1:wN,c1:wB,d1:wQ,e4:wP,e5:bP,f3:wN,c5:bB,c4:wB,d5:bP,g1:wK,f1:wR"
var fenFour string = "rnbqk1nr/ppp2ppp/8/2bpp3/2B1P3/5N2/PPPP1PPP/RNBQ1RK1 b kq - 1 4"

var tests = []testpair{
  {jsonStringOne, fenOne},
  {jsonStringTwo, fenTwo},
  {jsonStringThree, fenThree},
  {jsonStringFour, fenFour},
}

func TestPositions(t *testing.T) {
  for _, pair := range tests {
    b := GenerateBoard(pair.FEN)
    resMap := jsonToMap(pair.jsonPosition)
    if len(resMap) != len(b.whites) + len(b.blacks) {
      t.Error("Incorrect number of pieces. Expected ", len(resMap))
    }
    for pos, piece := range b.whites {
      if fenPieceToObj[piece] != resMap[pos.String()] {
        t.Error(fmt.Sprintf("Expected %s on %s, got %s", resMap[pos.String()], pos.String(), fenPieceToObj[piece]))
      }
    }
    for pos, piece := range b.blacks {
      if fenPieceToObj[piece] != resMap[pos.String()] {
        t.Error(fmt.Sprintf("Expected %s on %s, got %s", resMap[pos.String()], pos.String(), fenPieceToObj[piece]))
      }
    }
  }
}

func TestRestOfGenerateBoard(t *testing.T) {
  for _, pair := range tests {
    fields := strings.Split(pair.FEN, " ")
    b := GenerateBoard(pair.FEN)

    whiteToMove := (fields[1] == "w")
    availableCastles := fields[2]
    enPassantSquare := fields[3]

    // Test turn
    if b.whiteToMove != whiteToMove {
      t.Error(fmt.Sprintf("Expected %t, got %t", whiteToMove, b.whiteToMove))
    }
    // Test available castles
    count := 0
    for _, avail := range b.availableCastles {
      if avail {
        count += 1
      }
    }
    if count != len(availableCastles) {
      t.Error(fmt.Sprintf("Expected %d available castles, got %d", len(availableCastles), len(b.availableCastles)))
    }
    // Test set enpassant square
    if enPassantSquare == "-" && b.enPassantSquare.row != -1 {
      t.Error("Expected null position, got ", b.enPassantSquare)
    }
    if enPassantSquare != "-" {
      if b.enPassantSquare.col != strings.Index("abcdefgh", string(enPassantSquare[0])) {
        t.Error(fmt.Sprintf("Expected en passant square on %s col, got the %d'th", string(enPassantSquare[0]), b.enPassantSquare.col))
      }
      if b.enPassantSquare.row != int(enPassantSquare[1] - '0') - 1 {
        t.Error(fmt.Sprintf("Expected en passant square on %d row, got the %d'th", enPassantSquare[1], b.enPassantSquare.row))
      }
    }
  }
}
