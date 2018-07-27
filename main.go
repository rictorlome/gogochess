package main

import "fmt"

func main() {
	a, b := ToPos("a1"), ToPos("h8")
	fmt.Println("a-row", a.row, "a-col", a.col)
	fmt.Println("b-row", b.row, "b-col", b.col)
	startServer()
}
