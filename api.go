package main

import (
  "net/http"
  "strings"
  "fmt"
  "encoding/json"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  fmt.Println(message)
  fmt.Println(message)

  w.Write([]byte(message))
}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  for k, v := range r.Form {
    fmt.Println("key: ", k)
    fmt.Println("val: ", v)
    w.Write([]byte(k))
  }
}

type ManyFens struct {
  FenArr []string
}

type PieceMoves struct {
  Fen string
  Square string
  NextMoves []string
}


func GetFensEndpoint(w http.ResponseWriter, r *http.Request) {
  fens := ManyFens {
    []string {
      "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
      "rnbqkbnr/pppp1ppp/8/4p3/4P3/8/PPPP1PPP/RNBQKBNR w KQkq e6 0 2",
    },
  }
  json.NewEncoder(w).Encode(fens)
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

func GetNextMoves(w http.ResponseWriter, r *http.Request) {
  q := Queen{}
  q.isWhite = true
  pmt := []PieceMoves {
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "d5"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "c3"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "e4"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "b7"),
      GetPieceMoves("rnbqkbnr/ppp2ppp/3p4/3Qp3/4P3/1PN5/P1PP1PPP/R1B1KBNR w KQkq - 0 1", "f1"),
  }
  json.NewEncoder(w).Encode(pmt)
}

func startServer() {
  router := mux.NewRouter()
  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
  origins := handlers.AllowedOrigins([]string{"*"})
  router.HandleFunc("/", RootEndpoint).Methods("POST")
  router.HandleFunc("/GetFens", GetFensEndpoint).Methods("GET")
  router.HandleFunc("/GetMoves", GetNextMoves).Methods("GET")
  if err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)); err != nil {
    panic(err)
  }
}
