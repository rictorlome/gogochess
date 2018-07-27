package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var testBoards = []string{"rnbqk1nr/ppp2ppp/4p3/3p4/1b1P4/8/PPP1PPPP/RNBQKBNR w KQkq - 0 1"}

func startServer() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	router.HandleFunc("/tests", TestEndpoint).Methods("GET", "POST")
	if err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)); err != nil {
		panic(err)
	}
}

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(testBoards)
	case "POST":
		HandleTestPost(w, r)
	default:
		fmt.Println("OTHER")
	}
}

func HandleTestPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	switch r.Form["PostType"][0] {
	case "NextMoves":
		PostNextMoves(w, r)
	case "MovePiece":
		PostMovePiece(w, r)
	}
}

func PostNextMoves(w http.ResponseWriter, r *http.Request) {
	//params are fen, square, methodToTest, PostType
	r.ParseForm()
	NextMoves := GetPieceMoves(r.Form["fen"][0], r.Form["square"][0], r.Form["methodToTest"][0])
	json.NewEncoder(w).Encode(NextMoves)
}

func PostMovePiece(w http.ResponseWriter, r *http.Request) {
	//params are fen, uci, PostType
	r.ParseForm()
	b := GenerateBoard(r.Form["fen"][0])
	b.moveUCI(r.Form["uci"][0])
	ResultFen := b.GenerateFen()
	json.NewEncoder(w).Encode(ResultFen)
}

type MoveStatus struct {
	Move            string
	WouldCauseCheck bool
}

type PieceMoves struct {
	Fen    string
	Square string
	Moves  []MoveStatus
}

func GetPieceMoves(fen string, sq string, methodToTest string) PieceMoves {
	b := GenerateBoard(fen)
	pos := ToPos(sq)
	_, piece := b.findPiece(pos)
	moves := make(map[Position]bool)
	switch methodToTest {
	case "GetAttackingSquares":
		moves = piece.GetAttackingSquares(pos, &b, piece.GetDefaultMoveDiffs())
	case "GetPseudoLegalMoves":
		moves = piece.GetPseudoLegalMoves(pos, &b)
	case "GetLegalMoves":
		moves = piece.GetLegalMoves(pos, &b)
	case "GetAllLegalMoves":
		moves = b.GetAllLegalMoves(piece.IsWhite())
	}
	var sqs []MoveStatus
	for move, _ := range moves {
		sq := MoveStatus{move.String(), b.wouldCauseCheck(parseMove(sq + move.String()))}
		// sq := MoveStatus{move.String(), false}
		sqs = append(sqs, sq)
	}
	return PieceMoves{
		fen, sq, sqs,
	}
}
