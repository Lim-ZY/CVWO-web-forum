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
  postHandler := &posts.PostHandler{DB: db}

  return func(r chi.Router) {
    r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
      response, _ := users.HandleList(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Get("/", root.BasicHandler)
    r.Get("/topics", topicHandler.List)
    r.Post("/topics", topicHandler.Create)
    r.Put("/topics/{id}", topicHandler.UpdateByID)
    r.Delete("/topics/{id}", topicHandler.DeleteByID)
    r.Get("/topics/{id}", postHandler.List)
    r.Post("/topics/{id}", postHandler.Create)
    r.Put("/topics/{id}/{postID}", postHandler.UpdateByID)
    r.Delete("/topics/{id}/{postID}", postHandler.DeleteByID)
  }
}
