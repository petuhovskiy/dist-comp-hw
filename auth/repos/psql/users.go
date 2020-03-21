package psql

import (
	"auth/modeldb"
	"context"
	"github.com/jackc/pgx/v4"
)

type Users struct {
	conn *pgx.Conn
}

func NewUsers(conn *pgx.Conn) *Users {
	return &Users{
		conn: conn,
	}
}

func (r *Users) Migrate() error {
	_, err := r.conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS users(
			id				serial PRIMARY KEY,
			created_at  	timestamp DEFAULT current_timestamp NOT NULL,
			email       	text UNIQUE NOT NULL,
			password_hash 	bytea NOT NULL
		)`,
	)
	return err
}

func (r *Users) Create(p modeldb.User) (modeldb.User, error) {
	err := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO users(email, password_hash) VALUES ($1, $2) RETURNING id, created_at`,
		p.Email, p.PasswordHash,
	).Scan(&p.ID, &p.CreatedAt)

	return p, err
}

func (r *Users) FindByEmail(email string) (modeldb.User, error) {
	var p modeldb.User

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT
		id, created_at, email, password_hash
		FROM users
		WHERE email = $1`,
		email,
	).Scan(&p.ID, &p.CreatedAt, &p.Email, &p.PasswordHash)

	return p, err
}

func (r *Users) FindByID(id uint) (modeldb.User, error) {
	var p modeldb.User

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT
		id, created_at, email, password_hash
		FROM users
		WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.CreatedAt, &p.Email, &p.PasswordHash)

	return p, err
}
