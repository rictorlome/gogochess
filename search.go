package main

import (
	"fmt"
)

func (b *Board) isCapture(move Move) bool {
	capture, _ := b.findPiece(move.end)
	return capture
}

func (b *Board) isEnpassant(piece Piece, move Move) bool {
	return (piece.ToString() == "p" || piece.ToString() == "P") && move.end == b.enPassantSquare
}

func (b *Board) isCastle(piece Piece, move Move) bool {
	if piece.ToString() == "K" || piece.ToString() == "k" {
		return move.end.col > move.start.col+1 || move.end.col < move.start.col-1
	}
	return false
}

func (b *Board) isPromotion(piece Piece, move Move) bool {
	return piece.ToString() == "P" && move.end.row == 7 || piece.ToString() == "p" && move.end.row == 0
}

type perft struct {
	depth, nodes, captures, enpassants, castles, promotions, checks, checkmates int
}

func (p perft) String() string {
	return fmt.Sprintf("At depth %v,\n%v nodes, %v captures, %v enpassants, %v castles, %v promotions, %v checks, and %v checkmates", p.depth, p.nodes, p.captures, p.enpassants, p.castles, p.promotions, p.checks, p.checkmates)
}

type CurAndPrevBoards struct {
	curBoard Board
	prevBoard Board
}

type Node struct {
	b Board
	nextMoves []Move
	children []Node
}

func searchTree(initial Board, remainingDepth int) Node {
	var nextMoves []Move
	var resultChildren []Node
	if 0 < remainingDepth {
		nextMoves = initial.GetAllNextMoves(initial.whiteToMove)
		for _, nextMove := range nextMoves {
			newBoard := initial.Dup()
			newBoard.ApplyMove(nextMove)
			resultChildren = append(resultChildren, searchTree(*newBoard, remainingDepth-1))
		}
	}
	return Node{initial, nextMoves, resultChildren}
}

func walkTree(prefix string, n Node) {
	if len(n.nextMoves) == 0 {
		fmt.Println(prefix)
		return
	}
	for index, child := range(n.children) {
		newPrefix := fmt.Sprintf("%s %s", prefix, n.nextMoves[index].String())
		walkTree(newPrefix, child)
	}
}

func size(n Node) int {
	res := 1
	if 0 < len(n.nextMoves) {
		for _, child := range(n.children) {
			res += size(child)
		}
	}
	return res
}

func countLeaves(n Node) int {
	res := 0
	if 0 < len(n.nextMoves) {
		for _, child := range(n.children) {
			res += countLeaves(child)
		}
	} else {
		res = 1
	}
	return res
}

func divideTree(prefix string, n Node, curDepth int, maxDepth int, i int) {
	if maxDepth < curDepth {
		return
	}
	for idx, child := range(n.children) {
		newPrefix := fmt.Sprintf("%v %v", prefix, n.nextMoves[idx].String())
		divideTree(newPrefix, child, curDepth + 1, maxDepth, i)
	}
	if 1 < curDepth && curDepth < i {
		fmt.Println(fmt.Sprintf("%v. %v moves =        %v",  curDepth, prefix, countLeaves(n)))
	}
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func divide(initial Board, maxDepth int, dividor int) {
	for i := 1; i <= maxDepth; i++ {
		root := searchTree(initial, i)
		if 2 < i {
			divideTree("", root, 1, dividor, i)
		}
		fmt.Println(fmt.Sprintf("perft( %v)=          %--v", i, countLeaves(root)))
	}
}

func search(initial *Board, maxPly int, divide int) (perft, []CurAndPrevBoards) {
	if maxPly <= 0 {
		return perft{0, 1, 0, 0, 0, 0, 0, 0}, []CurAndPrevBoards{
			CurAndPrevBoards{*initial, Board{}},
		}
	}

	prevPerft, prevNodes := search(initial, maxPly-1, divide)
	var curPerft perft
	var curNodes []CurAndPrevBoards
	divideMoves := make(map[string]int)

	curPerft.depth = prevPerft.depth + 1
	for _, CurAndPrevBoard := range prevNodes {
		cur, prev := CurAndPrevBoard.curBoard, CurAndPrevBoard.prevBoard
		boardsNextMoves := cur.GetAllNextMoves(cur.whiteToMove)
		if divide == maxPly - 1 {
			divideMoves[prev.GenerateFen()] += len(boardsNextMoves)
		}
		for _, move := range boardsNextMoves {
			_, piece := cur.findPiece(move.start)
			if cur.isCapture(move) {
				curPerft.captures += 1
			}
			if cur.isEnpassant(piece, move) {
				curPerft.enpassants += 1
			}
			if cur.isCastle(piece, move) {
				curPerft.castles += 1
			}
			if cur.isPromotion(piece, move) {
				curPerft.promotions += 1
			}
			newBoard := cur.Dup()
			newBoard.ApplyMove(move)
			if newBoard.inCheck(true) || newBoard.inCheck(false) {
				curPerft.checks += 1
			}
			if newBoard.inCheckmate(true) || newBoard.inCheckmate(false) {
				curPerft.checkmates += 1
			}
			curAndPrev := CurAndPrevBoards{*newBoard, cur}
			curNodes = append(curNodes, curAndPrev)
		}
	}
	curPerft.nodes = len(curNodes)
	if divide == maxPly - 1 {
		for k, v := range divideMoves {
			fmt.Println(fmt.Sprintf("%v. %v moves =       %v",divide,k,v))
		}
		fmt.Println(curPerft.depth, curPerft.nodes)
	}
	return curPerft, curNodes
}
