package main

var MAX_SCORE int = 10000
var MIN_SCORE int = -10000

func Evaluate(b *Board) int {
	if b.inCheckmate(false) {
		return MAX_SCORE
	}
	if b.inCheckmate(true) {
		return MIN_SCORE
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
	evaluation := Evaluate(b)
	if (evaluation == MAX_SCORE && max) || (evaluation == MIN_SCORE && !max) {
		return evaluation, []Board{*b}
	}
	if maxDepth <= 0 {
		return evaluation, []Board{}
	}

	var Best int
	if max {
		Best = MIN_SCORE
	} else {
		Best = MAX_SCORE
	}
	var Bestpath []Board
	Boards := []Board{*b}

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
