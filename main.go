package main

import "fmt"

func main() {
	b := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fmt.Println(search(&b,0))
	// startServer()
}
