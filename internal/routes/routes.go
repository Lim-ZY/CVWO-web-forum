package routes

import (
  "os"

  pg "github.com/lim-zy/CVWO-web-forum/internal/database"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/users"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/topics"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/posts"
  "github.com/lim-zy/CVWO-web-forum/internal/handlers/postview"
  "github.com/go-chi/chi/v5"
)

func GetRoutes(db *pg.Database) func(r chi.Router) {
  topicHandler := &topics.TopicHandler{DB: db}
  postHandler := &posts.PostHandler{DB: db}
  viewHandler := &postview.ViewHandler{DB: db}
  userHandler := &users.UserHandler{DB: db, 
                                    CookieSecret: os.Getenv("JWT_SECRET"), 
                                    CookieName: os.Getenv("COOKIE_NAME")}

  return func(r chi.Router) {
    r.Post("/login", userHandler.Login)
    r.Get("/user", userHandler.GetUser)
    r.Post("/logout", userHandler.Logout)
    r.Get("/t", topicHandler.List)
    r.Get("/t{id}", topicHandler.GetByID)
    r.Post("/t", topicHandler.Create)
    r.Put("/t/{id}", topicHandler.UpdateByID)
    r.Delete("/t/{id}", topicHandler.DeleteByID)
    r.Get("/t/{id}", postHandler.List)
    //r.Get("/t{id}/{postID}", postHandler.GetByID)
    r.Post("/t/{id}", postHandler.Create)
    r.Put("/t/{id}/{postID}", postHandler.UpdateByID)
    r.Delete("/t/{id}/{postID}", postHandler.DeleteByID)
    r.Get("/t/{id}/{postID}", viewHandler.Get)
    r.Post("/t/{id}/{postID}", viewHandler.AddComment)
    r.Put("/t/{id}/{postID}/{commentID}", viewHandler.UpdateCommentByID)
    r.Delete("/t/{id}/{postID}/{commentID}", viewHandler.DeleteCommentByID)
  }
}
