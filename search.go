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

type BoardAndMove struct {
	board Board
	move  Move
}

func search(initial *Board, maxPly int, divide int) (perft, []BoardAndMove) {
	if maxPly <= 0 {
		return perft{0, 1, 0, 0, 0, 0, 0, 0}, []BoardAndMove{
			BoardAndMove{*initial, Move{}},
		}
	}

	prevPerft, prevNodes := search(initial, maxPly-1, divide)
	var curPerft perft
	var curNodes []BoardAndMove
	divideMoves := make(map[string]int)

	curPerft.depth = prevPerft.depth + 1
	for _, boardAndMove := range prevNodes {
		b := boardAndMove.board
		boardsNextMoves := b.GetAllNextMoves(b.whiteToMove)
		if divide == maxPly-1 {
			divideMoves[boardAndMove.move.String()] += len(boardsNextMoves)
		}
		for _, move := range boardsNextMoves {
			if b.isCapture(move) {
				curPerft.captures += 1
			}
			if b.isEnpassant(move) {
				curPerft.enpassants += 1
			}
			if b.isCastle(move) {
				curPerft.castles += 1
			}
			if b.isPromotion(move) {
				curPerft.promotions += 1
			}
			newBoard := b.Dup()
			newBoard.ApplyMove(move)
			if newBoard.inCheck(true) || newBoard.inCheck(false) {
				curPerft.checks += 1
			}
			if newBoard.inCheckmate(true) || newBoard.inCheckmate(false) {
				curPerft.checkmates += 1
			}
			bAndM := BoardAndMove{*newBoard, move}
			curNodes = append(curNodes, bAndM)
		}
	}
	curPerft.nodes = len(curNodes)
	if divide == maxPly-1 {
		for k, v := range divideMoves {
			fmt.Println(divide, ".", k, v)
		}
		fmt.Println(curPerft.depth, curPerft.nodes)
	}
	return curPerft, curNodes
}
