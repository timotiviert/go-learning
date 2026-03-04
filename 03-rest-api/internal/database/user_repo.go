package database

import (
	"context"

	"github.com/timotiviert/go-learning/03-rest-api/internal/models"
)

// Interface for decoupling :)) -> can test with mocks and exchange Db, if necessary.
type UserRepository interface {
	Create(email, username, hashedPassword string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
}

type postgresUserRepository struct {
	db *DB
}

func NewUserRepository(db *DB) UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func (r *postgresUserRepository) Create(email, username, hashedPassword string) (*models.User, error) {
	var u models.User

	err := r.db.Pool.QueryRow(context.Background(),
		"INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id, email, username",
		email, username, hashedPassword,
	).Scan(&u.ID, &u.Email, &u.Username)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *postgresUserRepository) GetByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.Pool.QueryRow(context.Background(),
		"SELECT id, email, username, created_at, updated_at FROM users WHERE email = $1",
		email,
	).Scan(&u.ID, &u.Email, &u.Username, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *postgresUserRepository) GetByUsername(username string) (*models.User, error) {
	var u models.User
	err := r.db.Pool.QueryRow(context.Background(),
		"SELECT id, email, username, created_at, updated_at FROM users WHERE username = $1",
		username,
	).Scan(&u.ID, &u.Email, &u.Username, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
