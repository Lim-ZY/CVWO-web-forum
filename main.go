package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  //"os/signal"
  "context"
  "github.com/lim-zy/CVWO-web-forum/internal/router"
  "github.com/lim-zy/CVWO-web-forum/internal/database"
  "github.com/joho/godotenv"
)

func main() {
  //ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
  //defer cancel()
  err := godotenv.Load()
  if err != nil {
    log.Fatalln("Error loading .env file")
  }

  db, err := database.GetDB(context.Background())
  if err != nil {
    log.Fatalln("Error connecting to PostgreSQL Database. Quitting...")
  }

  r := router.Setup(db)
  port := os.Getenv("PORT")
  if port == "" {
    port = "8000"
  }
  fmt.Printf("Listening on port %v at http://localhost:%v ...\n", port, port)
  addr := ":" + port
  log.Fatalln(http.ListenAndServe(addr, r))
}
