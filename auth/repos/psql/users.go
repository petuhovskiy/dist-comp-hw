package psql

import (
	"auth/modeldb"
	"context"
	"github.com/jackc/pgx/v4"
	"lib/pb"
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
			email       	text UNIQUE,
			phone       	text UNIQUE,
			password_hash 	bytea NOT NULL,
			role			integer NOT NULL
		)`,
	)
	return err
}

func (r *Users) Create(p modeldb.User) (modeldb.User, error) {
	var email, phone interface{}
	if p.Email != "" {
		email = p.Email
	}
	if p.Phone != "" {
		phone = p.Phone
	}

	err := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO users(email, phone, password_hash, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		email, phone, p.PasswordHash, p.Role,
	).Scan(&p.ID, &p.CreatedAt)

	return p, err
}

func (r *Users) FindByLogin(login string) (modeldb.User, error) {
	var p modeldb.User

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT
		id, created_at, email, phone, password_hash, role
		FROM users
		WHERE email = $1 OR phone = $1`,
		login,
	).Scan(&p.ID, &p.CreatedAt, &p.Email, &p.Phone, &p.PasswordHash, &p.Role)

	return p, err
}

func (r *Users) FindByID(id uint) (modeldb.User, error) {
	var p modeldb.User

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT
		id, created_at, email, login, password_hash, role
		FROM users
		WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.CreatedAt, &p.Email, &p.Phone, &p.PasswordHash, &p.Role)

	return p, err
}

func (r *Users) UpdateEmail(id uint, email string) error {
	_, err := r.conn.Exec(
		context.Background(),
		`UPDATE users
		SET email = $2
		WHERE id = $1`,
		id,
		email,
	)
	return err
}

func (r *Users) UpdatePhone(id uint, phone string) error {
	_, err := r.conn.Exec(
		context.Background(),
		`UPDATE users
		SET phone = $2
		WHERE id = $1`,
		id,
		phone,
	)
	return err
}

func (r *Users) UpdateRole(id uint, role pb.AuthRole) error {
	_, err := r.conn.Exec(
		context.Background(),
		`UPDATE users
		SET role = $2
		WHERE id = $1`,
		id,
		role,
	)
	return err
}
