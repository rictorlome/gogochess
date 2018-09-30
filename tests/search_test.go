package main

import (
  "testing"
  "fmt"
)

type perftSeq struct {
  startFen string
  perfts []perft
}
//tables taken from https://chessprogramming.wikispaces.com/Perft%20Results
//note no castles or promotions
var initialPerft = perftSeq {
  "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
  []perft{
    perft{0,1,0,0,0,0,0,0},
    perft{1,20,0,0,0,0,0,0},
    perft{2,400,0,0,0,0,0,0},
    perft{3,8902,34,0,0,0,12,0},
    perft{4,197281,1576,0,0,0,469,8},
    perft{5,4865609,82719,258,0,0,27351,347},
    perft{6,119060324,2812008,5248,0,0,809099,10823},
  },
}

var secondaryPerf = perftSeq {
  "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1",
  []perft{
    perft{0,1,0,0,0,0,0,0},
    perft{1,48,8,0,2,0,0,0},
    perft{2,2039,351,1,91,0,3,0},
    perft{3,97862,17102,45,3162,0,993,1},
    perft{4,4085603,757163,1929,128013,15172,25523,43},
    perft{5,193690690,35043416,73365,4993637,8392,3309887,30171},
  },
}

var perftTests = []perftSeq{
  initialPerft, secondaryPerf,
}



func TestSearch(t *testing.T) {
  for _, pSeq := range perftTests {
    b := GenerateBoard(pSeq.startFen)
    for i := 0; i < 5; i++ {
      perft, _ := search(&b, i)
      if perft != pSeq.perfts[i] {
        t.Error(fmt.Sprintf("\nExpected:\n%v\nGot:\n%v", pSeq.perfts[i], perft))
      }
    }
  }
}
