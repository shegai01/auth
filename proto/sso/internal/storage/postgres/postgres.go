package postgres

import (
	"auth/internal/domain/models"
	"auth/storage"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type Storage struct {
	DB *sql.DB
}

func New(dsn string) (*Storage, error) {
	const (
		op = "storage.postgres.New"
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const (
		op = "storage.Postgres.SaveUser"
	)

	stmt, err := s.DB.Prepare("INSERT INTO users(email, pass_hash) values (?, ? )")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, email, passHash)
	if err != nil {
		var (
			err         error
			postgresErr pq.Error
		)
		if errors.As(err, &postgresErr) && postgresErr.Code == "23505" {

			return 0, fmt.Errorf("%s: %w", op, err)
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)

	}
	return id, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const (
		op = "storage.Postgres.User"
	)
	var (
		user models.User
	)
	stmt, err := s.DB.Prepare("select id, email, pass_hash from users where email=?")
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRowContext(ctx, email)
	err = row.Scan(&user.ID, &user.Email, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}
	return user, nil
}
