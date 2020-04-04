package psql

import (
	"auth/modeldb"
	"context"
	"github.com/jackc/pgx/v4"
)

type Confirm struct {
	conn *pgx.Conn
}

func NewConfirm(conn *pgx.Conn) *Confirm {
	return &Confirm{
		conn: conn,
	}
}

func (r *Confirm) Migrate() error {
	_, err := r.conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS confirms(
			link		text PRIMARY KEY,
			user_id     integer NOT NULL REFERENCES users(id),
			type		text NOT NULL,
			subject		text NOT NULL,
			expire_at   timestamp NOT NULL
		)`,
	)
	return err
}

func (r *Confirm) Create(p modeldb.Confirm) (modeldb.Confirm, error) {
	_, err := r.conn.Exec(
		context.Background(),
		`INSERT INTO confirms(link, user_id, type, subject, expire_at) VALUES ($1, $2, $3, $4, $5)`,
		p.Link, p.UserID, p.Type, p.Subject, p.ExpireAt,
	)

	return p, err
}

func (r *Confirm) Find(link string) (modeldb.Confirm, error) {
	var p modeldb.Confirm

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT
		link, user_id, type, subject, expire_at
		FROM confirms
		WHERE link = $1 AND expire_at > current_timestamp`,
		link,
	).Scan(&p.Link, &p.UserID, &p.Type, &p.Subject, &p.ExpireAt)

	return p, err
}

func (r *Confirm) Delete(link string) error {
	_, err := r.conn.Exec(
		context.Background(),
		`DELETE FROM confirms WHERE link=$1`,
		link,
	)

	return err
}
