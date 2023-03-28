package main

import (
	"log"
	"net/http"
  "database/sql"
  "github.com/lib/pq"
)


const (
  host = 'localhost'
  port = 8001
  user = 'postgres'
  password = 'passwd'
  dbname = 'postgres'
)

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
  "password=%s dbname=%s sslmode=disable",
  host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  
  err = db.Ping()
  if err != nil {
    panic(err)
  }
  
  fmt.Println("db connected")

  router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
  
  fmt.Println("api up")
