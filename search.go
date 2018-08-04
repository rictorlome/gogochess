package main

import "fmt"

func (b *Board) isCapture(move Move) bool {
	capture, _ := b.findPiece(move.end)
	return capture
}

func (b *Board) isEnpassant(move Move) bool {
	_, piece := b.findPiece(move.start)
	return (piece.ToString() == "p" || piece.ToString() == "P") && move.end == b.enPassantSquare
}

func (b *Board) isCastle(move Move) bool {
	_, piece := b.findPiece(move.start)
	if piece.ToString() == "K" || piece.ToString() == "k" {
		return move.end.col > move.start.col+1 || move.end.col < move.start.col-1
	}
	return false
}

func (b *Board) isPromotion(move Move) bool {
	_, piece := b.findPiece(move.start)
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
			if cur.isCapture(move) {
				curPerft.captures += 1
			}
			if cur.isEnpassant(move) {
				curPerft.enpassants += 1
			}
			if cur.isCastle(move) {
				curPerft.castles += 1
			}
			if cur.isPromotion(move) {
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
