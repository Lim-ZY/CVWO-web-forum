package router

import (
  "github.com/lim-zy/CVWO-web-forum/internal/routes"
  pg "github.com/lim-zy/CVWO-web-forum/internal/database"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"
)

func Setup(db *pg.Database) chi.Router {
  r := chi.NewRouter()
  r.Use(cors.Handler(cors.Options{
    AllowedOrigins:   []string{"http://localhost:3000/*"}, // Use this to allow specific origin hosts
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

  setUpRoutes(r, db)
  return r
}

func setUpRoutes(r chi.Router, db *pg.Database) {
  r.Group(routes.GetRoutes(db))
}
