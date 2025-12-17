package main

import (
  "fmt"
  "log"
  "net/http"
  //"os"
  //"os/signal"
  "context"
  "github.com/lim-zy/CVWO-web-forum/internal/router"
  "github.com/lim-zy/CVWO-web-forum/internal/database"
)

func main() {
  //ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
  //defer cancel()

  db, err := database.GetDB(context.Background())
  if err != nil {
    log.Fatalln("Error connecting to PostgreSQL Database. Quitting...")
  }

  r := router.Setup(db)
  fmt.Print("Listening on port 8000 at http://localhost:8000 ...")
  log.Fatalln(http.ListenAndServe(":8000", r))
}
