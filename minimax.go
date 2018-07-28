package main

func Evaluate(b *Board) int {
	if b.inCheckmate(false) {
		return 1
	}
	if b.inCheckmate(true) {
		return -1
	}
	return 0
}

func GetSuccessors(b *Board) []Board {
	var Boards []Board
	nextMoves := b.GetAllNextMoves(b.whiteToMove)
	for _, move := range nextMoves {
		newBoard := b.Dup()
		newBoard.ApplyMove(move)
		Boards = append(Boards, *newBoard)
	}
	return Boards
}

func Minimax(b *Board, maxDepth int, max bool) (int, []Board) {
	if maxDepth <= 1 {
		return Evaluate(b), []Board{}
	}
	var Best int
	var Bestpath []Board
	var Boards []Board
	Boards = append(Boards, *b)

	successors := GetSuccessors(b)

	for i, successor := range successors {
		optimalValue, path := Minimax(&successor, maxDepth-1, !max)
		if i == 0 || (max && optimalValue > Best) || (!max && optimalValue < Best) {
			Best = optimalValue
			Bestpath = path
		}
	}

	return Best, append(Boards, Bestpath...)
}
