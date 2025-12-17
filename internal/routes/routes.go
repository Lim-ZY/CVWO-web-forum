package routes

import (
  "encoding/json"
  "net/http"
  pg "github.com/lim-zy/CVWO-web-forum/internal/database"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/users"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/root"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/topics"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/posts"
  //"github.com/lim-zy/CVWO-web-forum/internal/handlers/post"
  "github.com/go-chi/chi/v5"
)

func GetRoutes(db *pg.Database) func(r chi.Router) {
  topicHandler := &topics.TopicHandler{DB: db}

  return func(r chi.Router) {
    r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
      response, _ := users.HandleList(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Get("/", root.BasicHandler)
    r.Get("/topics", func(w http.ResponseWriter, req *http.Request) {
      response, _ := topicHandler.List(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Post("/topics", func(w http.ResponseWriter, req *http.Request) {
      response, _ := topicHandler.Create(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Put("/topics/{id}", func(w http.ResponseWriter, req *http.Request) {
      response, _ := topicHandler.UpdateByID(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Delete("/topics/{id}", func(w http.ResponseWriter, req *http.Request) {
      topicHandler.DeleteByID(w, req)
      w.Header().Set("Content-Type", "application/json")
    })
    r.Get("/topics/{id}", posts.List)
    r.Post("/topics/{id}", posts.Create)
    //r.Put("/topics/{id}", posts.UpdateByID)
    //r.Delete("/topics/{id}", posts.DeleteByID)
  }
}
