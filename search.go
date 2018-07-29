package main

func isGoal(b *Board) bool {
	return b.inCheckmate(false)
}

func getSuccessors(b *Board) []Board {
	var Boards []Board
	nextMoves := b.GetAllNextMoves(b.whiteToMove)
	for _, move := range nextMoves {
		newBoard := b.Dup()
		newBoard.ApplyMove(move)
		Boards = append(Boards, *newBoard)
	}
	return Boards
}

type perft struct {
  depth, nodes, captures, enpassants, castles, promotions, checks, checkmates int
}

func search(initial *Board, maxPly int) (perft, []Board) {
	if maxPly <= 0 {
		return perft{0,1,0,0,0,0,0,0}, []Board{*initial}
	}

	prevPerft, prevNodes := search(initial, maxPly - 1)
	var curPerft perft
	var curNodes []Board

	curPerft.depth = prevPerft.depth + 1
	for _, board := range prevNodes {
		boardsNextMoves := board.GetAllNextMoves(board.whiteToMove)
		for _, move := range boardsNextMoves {
			// if board.isCapture(move) {
			// 	curPerft.captures += 1
			// }
			// if board.isEnpassant(move) {
			// 	curPerft.enpassants += 1
			// }
			// if board.isCastle(move) {
			// 	curPerft.castles += 1
			// }
			// if board.isPromotion(move) {
			// 	curPerft.promotions += 1
			// }
			newBoard := board.Dup()
			newBoard.ApplyMove(move)
			if newBoard.inCheck(true) || newBoard.inCheck(false) {
				curPerft.checks += 1
			}
			if newBoard.inCheckmate(true) || newBoard.inCheckmate(false) {
				curPerft.checkmates += 1
			}
			curNodes = append(curNodes,*newBoard)
		}
	}
	curPerft.nodes = len(curNodes)
	return curPerft, curNodes
}
