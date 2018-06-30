package main

import (
  "time"
  "fmt"
)

func main() {
  // fen := "rnbqk1nr/ppp2ppp/8/2bpp3/2B1P3/5N2/PPPP1PPP/RNBQ1RK1 b kq - 1 4"
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  start := time.Now()
  b := GenerateBoard(fen)
  // n := Knight{}
  // n.isWhite = true
  bish := Queen{}
  bish.isWhite = true
  res := bish.GetPseudoLegalMoves(Position{5,5}, &b)
  // // fmt.Println(n.isWhite)
  // res := n.GetPseudoLegalMoves(Position{0,0}, &b)
  fmt.Println(Position{5,5}.getNeighbors())
  fmt.Println(res)


  t := time.Now()
  fmt.Println(t.Sub(start))

  startServer()

}
