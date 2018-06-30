package main

import (
  "net/http"
  "strings"
  "fmt"

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

func startServer() {
  router := mux.NewRouter()
  headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
  methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
  origins := handlers.AllowedOrigins([]string{"*"})
  router.HandleFunc("/", RootEndpoint).Methods("POST")
  if err := http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)); err != nil {
    panic(err)
  }
}
