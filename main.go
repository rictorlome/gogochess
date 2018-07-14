package main

import (
  "time"
  "fmt"
)

func main() {
  // fen := "rnbqk1nr/ppp2ppp/8/2bpp3/2B1P3/5N2/PPPP1PPP/RNBQ1RK1 b kq - 1 4"
  // fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  fen := "rnbqk1nr/ppp2ppp/4p3/3p4/1b1P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1"
  start := time.Now()
  b := GenerateBoard(fen)
  fmt.Println(b.inCheck(true))
  fmt.Println(b.inCheck(false))
  startServer()


  t := time.Now()
  fmt.Println(t.Sub(start))
}
