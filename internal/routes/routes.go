package routes

import (
  "encoding/json"
  "net/http"
  pg "github.com/lim-zy/CVWO-web-forum/internal/database"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/users"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/root"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/topics"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/posts"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/postview"
  "github.com/go-chi/chi/v5"
)

func GetRoutes(db *pg.Database) func(r chi.Router) {
  topicHandler := &topics.TopicHandler{DB: db}
  postHandler := &posts.PostHandler{DB: db}
  viewHandler := &postview.ViewHandler{DB: db}

  return func(r chi.Router) {
    r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
      response, _ := users.HandleList(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Get("/", root.BasicHandler)
    r.Get("/t", topicHandler.List)
    r.Post("/t", topicHandler.Create)
    r.Put("/t/{id}", topicHandler.UpdateByID)
    r.Delete("/t/{id}", topicHandler.DeleteByID)
    r.Get("/t/{id}", postHandler.List)
    r.Post("/t/{id}", postHandler.Create)
    r.Put("/t/{id}/{postID}", postHandler.UpdateByID)
    r.Delete("/t/{id}/{postID}", postHandler.DeleteByID)
    r.Get("/t/{id}/{postID}", viewHandler.Get)
    r.Post("/t/{id}/{postID}", viewHandler.AddComment)
    r.Put("/t/{id}/{postID}/{commentID}", viewHandler.UpdateCommentByID)
    r.Delete("/t/{id}/{postID}/{commentID}", viewHandler.DeleteCommentByID)
  }
}
