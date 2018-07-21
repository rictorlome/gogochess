package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  // "strconv"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

var initialBoard string = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func startServer() {
  router := mux.NewRouter()
  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
  origins := handlers.AllowedOrigins([]string{"*"})
  router.HandleFunc("/tests", TestEndpoint).Methods("GET","POST")
  router.HandleFunc("/GetMoves", GetNextMoves).Methods("GET")
  if err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)); err != nil {
    panic(err)
  }
}

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    json.NewEncoder(w).Encode(initialBoard)
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
    PostNextMoves(w,r)
  case "MovePiece":
    PostMovePiece(w,r)
  }
}

func PostNextMoves(w http.ResponseWriter, r *http.Request) {
  //params are fen, square, PostType
  r.ParseForm()
  NextMoves := GetPieceMoves(r.Form["fen"][0], r.Form["square"][0])
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
  Move string
  WouldCauseCheck bool
}

type PieceMoves struct {
  Fen string
  Square string
  Moves []MoveStatus
}

 func GetNextMoves(w http.ResponseWriter, r *http.Request) {
  q := Queen{}
  q.isWhite = true
  pmt := []PieceMoves {
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "d5"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "c3"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "e4"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "b7"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "f1"),
      GetPieceMoves("rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 1", "e5"),
      GetPieceMoves("rnb1kb1r/ppp1p1pp/3q1n2/3pPp2/3P4/8/PPP2PPP/RNBQKBNR w KQkq - 0 1", "e5"),
      // GetPieceMoves("r1bqk2r/pppp1ppp/2n2n2/2b1p3/2B1P3/2P2N2/PP1P1PPP/RNBQK2R w KQkq - 0 1", "e1"),
  }
  json.NewEncoder(w).Encode(pmt)
}

func GetPieceMoves(fen string, sq string) PieceMoves {
  b := GenerateBoard(fen)
  pos := ToPos(sq)
  _, piece := b.findPiece(pos)
  moves := piece.GetPseudoLegalMoves(pos, &b)
  var sqs []MoveStatus
  for move, _ := range(moves) {
    sq := MoveStatus{move.String(), b.wouldCauseCheck(parseMove(sq+move.String()))}
    sqs = append(sqs, sq)
  }
  return PieceMoves {
    fen, sq, sqs,
  }
}
