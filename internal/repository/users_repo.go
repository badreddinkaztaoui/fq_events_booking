package repository

import (
	"context"
	"database/sql"

	"github.com/badreddinkaztaoui/fq_events_booking/internal/models"
)

type UsersRepo struct {
	DB *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{DB: db}
}

func (repo *UsersRepo) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users(first_name, last_name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := repo.DB.QueryRowContext(ctx, query, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	return row.Scan(&user.ID)
}

func (repo *UsersRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, email, password_hash
		FROM users
		WHERE email = $1
	`
	row := repo.DB.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
