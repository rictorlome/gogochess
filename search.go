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

func search(initial *Board, maxDepth int) (bool, []Board) {
	var path []Board

	if maxDepth <= 0 {
		return false, path
	}

	path = append(path, *initial)

	if isGoal(initial) {
		return true, path
	}

	successors := getSuccessors(initial)

	for _, board := range successors {
		reachedGoal, childPath := search(&board, maxDepth-1)
		if reachedGoal {
			path = append(path, childPath...)
			return true, path
		}
	}
	return false, []Board{}
}
