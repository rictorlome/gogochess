package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "strconv"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  for k, v := range r.Form {
    fmt.Println("key: ", k)
    fmt.Println("val: ", v)
    w.Write([]byte(k))
  }
}

var fens = []string{"rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1",
      "rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1",
      "rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1",
      "rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1",
      "rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1",
      "rnbqkbnr/ppp1p1pp/8/3pPp2/8/8/PPPP1PPP/RNBQKBNR w KQkq f6 0 1",
      "rnb1kb1r/ppp1p1pp/3q1n2/3pPp2/3P4/8/PPP2PPP/RNBQKBNR w KQkq - 0 1",}

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
  case "GET":
    json.NewEncoder(w).Encode(fens)
  case "POST":
    PostNextMoves(w, r)
  default:
    fmt.Println("OTHER")
  }
}

type PieceMoves struct {
  Fen string
  Square string
  NextMoves []string
}

func GetPieceMoves(fen string, sq string) PieceMoves {
  b := GenerateBoard(fen)
  pos := ToPos(sq)
  _, piece := b.findPiece(pos)
  moves := piece.GetPseudoLegalMoves(pos, &b)
  sqs := []string{}
  for _, move := range(moves) {
    sqs = append(sqs, move.String())
  }
  return PieceMoves {
    fen, sq, sqs,
  }
}

func PostNextMoves(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  id, _ := strconv.Atoi(r.Form["id"][0])
  requestedFen := fens[id]
  NextMoves := GetPieceMoves(requestedFen, r.Form["square"][0])
  json.NewEncoder(w).Encode(NextMoves)
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

func startServer() {
  router := mux.NewRouter()
  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
  origins := handlers.AllowedOrigins([]string{"*"})
  router.HandleFunc("/", RootEndpoint).Methods("POST")
  router.HandleFunc("/tests", TestEndpoint).Methods("GET","POST")
  router.HandleFunc("/GetMoves", GetNextMoves).Methods("GET")
  if err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)); err != nil {
    panic(err)
  }
}
