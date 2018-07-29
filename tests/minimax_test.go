package main

import (
  "testing"
  "fmt"
)

type MateInTwoTest struct {
  initial string
  solution string
}
// from https://sites.google.com/site/darktemplarchess/mate-in-2-puzzles
func TestWhiteToMateInTwo(t *testing.T) {
  tests := []MateInTwoTest{
    MateInTwoTest{
      "2bqkbn1/2pppp2/np2N3/r3P1p1/p2N2B1/5Q2/PPPPKPP1/RNB2r2 w KQkq - 0 1",
      "2bqkbn1/2pppQ2/np2N3/r3P1p1/p2N2B1/8/PPPPKPP1/RNB2r2 b KQkq - 0 1",
    },
    MateInTwoTest{
      "8/6K1/1p1B1RB1/8/2Q5/2n1kP1N/3b4/4n3 w - - 0 1",
      "8/6K1/1p3RB1/8/2Q5/B1n1kP1N/3b4/4n3 b - - 1 1",
    },
    MateInTwoTest{
      "B7/K1B1p1Q1/5r2/7p/1P1kp1bR/3P3R/1P1NP3/2n5 w - - 0 1",
      "8/K1B1p1Q1/2B2r2/7p/1P1kp1bR/3P3R/1P1NP3/2n5 b - - 1 1",
    },
  }
  for _, test := range tests {
    b := GenerateBoard(test.initial)
    _, boards := Minimax(&b, 4)
    result := boards[1].GenerateFen()
    if result != test.solution {
      t.Error(fmt.Sprintf("Expected %v for solution to %v, got %v", test.solution, test.initial, result))
    }
  }
}
