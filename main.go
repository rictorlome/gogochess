package main

import (
	// "os"
	"fmt"
	// "strconv"
)

// import "fmt"

func main() {
	// b := GenerateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")


	// c := *b.Dup()
	// pb := 0
	// pc := 0
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(fmt.Sprintf("dups are same?? %v", b.GenerateFen() == c.GenerateFen()))
	// 	bt := searchTree(b, 2)
	// 	ct := searchTree(c, 2)
	//
	// 	sb := size(bt)
	// 	sc := size(ct)
	//
	// 	fmt.Println(fmt.Sprintf("b size is %v, prev size is %v, equal? %v", sb, pb, sb == pb))
	// 	fmt.Println(fmt.Sprintf("c size is %v, prev size is %v, equal? %v", sc, pc, sc == pc))
	// 	fmt.Println(fmt.Sprintf("b and c equal? %v, b and c prev equal? %v", sb == sc, pb == pc))
	// 	fmt.Println("")
	//
	// 	pb = sb
	// 	pc = sc
	// }


	// node := searchTree(b, 3)
	// i, _ := strconv.Atoi(os.Args[1])
	// e, _ := strconv.Atoi(os.Args[2])
	m := make(map[Position]int)
	for k := 0; k < 1000; k++ {
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				m[Position{i,j}] += 1
			}
		}
	}
	fmt.Println(m)

	// startServer()
}
