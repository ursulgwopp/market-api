package repository

import (
	"sort"

	"github.com/jmoiron/sqlx"
	"github.com/ursulgwopp/go-market-app/internal/models"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetUserByID(id int) (models.User, error) {
	var user models.User
	query := `SELECT username, email FROM users WHERE id = $1`
	err := r.db.Get(&user, query, id)
	user.Id = id

	return user, err
}

func (r *UserPostgres) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, username, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var user models.User

		if err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	sortById := func(i, j int) bool {
		return users[i].Id < users[j].Id
	}
	sort.Slice(users, sortById)

	return users, nil
}