package users

import (
	"github.com/lim-zy/CVWO-web-forum/internal/database"
	"github.com/lim-zy/CVWO-web-forum/internal/models"
)

func List(db *database.Database) ([]models.User, error) {
	users := []models.User{
		{
			ID:   1,
			Name: "CVWO",
		},
    {
      ID:   2,
      Name: "cool",
    },
	}
	return users, nil
}
