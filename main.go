package main

import "fmt"

func main() {
	// b := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	b := GenerateBoard("8/6r1/8/1R6/8/2R5/8/8 w KQkq - 0 1")
	for i := 0; i < 10; i++ {
		fmt.Println(countLeaves(searchTree(b, 4)))
	}
	// startServer()
}
