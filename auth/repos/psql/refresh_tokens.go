package psql

import (
	"auth/modeldb"
	"context"
	"github.com/jackc/pgx/v4"
)

type RefreshTokens struct {
	conn *pgx.Conn
}

func NewRefreshTokens(conn *pgx.Conn) *RefreshTokens {
	return &RefreshTokens{
		conn: conn,
	}
}

func (r *RefreshTokens) Migrate() error {
	_, err := r.conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS refresh_tokens(
			id			serial PRIMARY KEY,
			created_at  timestamp DEFAULT current_timestamp NOT NULL,
			user_id     integer NOT NULL REFERENCES users(id),
			token		text UNIQUE NOT NULL,
			expire_at   timestamp NOT NULL
		)`,
	)
	return err
}

func (r *RefreshTokens) Create(p modeldb.RefreshToken) (modeldb.RefreshToken, error) {
	err := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO refresh_tokens(token, user_id, expire_at) VALUES ($1, $2, $3) RETURNING id, created_at`,
		p.Token, p.UserID, p.ExpireAt,
	).Scan(&p.ID, &p.CreatedAt)

	return p, err
}

func (r *RefreshTokens) FindNonExpired(token string) (modeldb.RefreshToken, error) {
	var p modeldb.RefreshToken

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT
		id, created_at, token, user_id, expire_at
		FROM refresh_tokens
		WHERE token = $1 AND expire_at > current_timestamp`,
		token,
	).Scan(&p.ID, &p.CreatedAt, &p.Token, &p.UserID, &p.ExpireAt)

	return p, err
}

func (r *RefreshTokens) Delete(id uint) error {
	_, err := r.conn.Exec(
		context.Background(),
		`DELETE FROM refresh_tokens WHERE id=$1`,
		id,
	)

	return err
}
