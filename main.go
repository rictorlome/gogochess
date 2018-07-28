package main

import "fmt"

func main() {
	b := GenerateBoard("2bqkbn1/2pppp2/np2N3/r3P1p1/p2N2B1/5Q2/PPPPKPP1/RNB2r2 w KQkq - 0 1")
	fmt.Println(Minimax(&b, 4, true))
	startServer()
}
