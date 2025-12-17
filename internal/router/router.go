package router

import (
  "github.com/lim-zy/CVWO-web-forum/internal/routes"
  pg "github.com/lim-zy/CVWO-web-forum/internal/database"
  "github.com/go-chi/chi/v5"
)

func Setup(db *pg.Database) chi.Router {
  r := chi.NewRouter()
  setUpRoutes(r, db)
  return r
}

func setUpRoutes(r chi.Router, db *pg.Database) {
  r.Group(routes.GetRoutes(db))
}
