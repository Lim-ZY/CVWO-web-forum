package routes

import (
  "encoding/json"
  "net/http"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/users"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/root"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/topics"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/posts"
  //"github.com/lim-zy/CVWO-web-forum/internal/handlers/post"
  "github.com/go-chi/chi/v5"
)

func GetRoutes() func(r chi.Router) {
  return func(r chi.Router) {
    r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
      response, _ := users.HandleList(w, req)
      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    })
    r.Get("/", root.BasicHandler)
    r.Get("/topics", topics.List)
    r.Post("/topics", topics.Create)
    //r.Put("/topics/{id}", topics.UpdateByID)
    //r.Delete("/topics/{id}", topics.DeleteByID)
    r.Get("/topics/{id}", posts.List)
    r.Post("/topics/{id}", posts.Create)
    //r.Put("/topics/{id}", posts.UpdateByID)
    //r.Delete("/topics/{id}", posts.DeleteByID)
  }
}
