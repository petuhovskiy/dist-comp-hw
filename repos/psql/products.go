package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/petuhovskiy/dist-comp-hw/modeldb"
)

type Products struct {
	conn *pgx.Conn
}

func NewProducts(conn *pgx.Conn) *Products {
	return &Products{
		conn: conn,
	}
}

func (r *Products) Migrate() error {
	_, err := r.conn.Exec(
		context.Background(),
		`CREATE TABLE IF NOT EXISTS products(
			id			serial PRIMARY KEY,
			name		text,
			code		text,
			category	text
		)`,
	)
	return err
}

func (r *Products) Save(p modeldb.Product) (modeldb.Product, error) {
	if p.ID != 0 {
		return r.Update(p)
	}

	err := r.conn.QueryRow(
		context.Background(),
		`INSERT INTO products(name, code, category) VALUES ($1, $2, $3) RETURNING id`,
		p.Name, p.Code, p.Category,
	).Scan(&p.ID)

	return p, err
}

func (r *Products) Update(p modeldb.Product) (modeldb.Product, error) {
	_, err := r.conn.Exec(
		context.Background(),
		`UPDATE products SET name=$2, code=$3, category=$4 WHERE id=$1`,
		p.ID, p.Name, p.Code, p.Category,
	)

	return p, err
}

func (r *Products) Get(id uint) (modeldb.Product, error) {
	var p modeldb.Product

	err := r.conn.QueryRow(
		context.Background(),
		`SELECT id, name, code, category FROM products WHERE id=$1`,
		id,
	).Scan(&p.ID, &p.Name, &p.Code, &p.Category)

	return p, err
}

func (r *Products) ListPage(limit, offset uint) ([]modeldb.Product, error) {
	query := `SELECT id, name, code, category FROM products ORDER BY id DESC`

	if limit != 0 {
		query += fmt.Sprintf(` LIMIT %d`, limit)
	}
	if offset != 0 {
		query += fmt.Sprintf(` OFFSET %d`, offset)
	}

	rows, err := r.conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var list []modeldb.Product
	for rows.Next() {
		var p modeldb.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Code, &p.Category)
		if err != nil {
			return nil, err
		}

		list = append(list, p)
	}

	return list, nil
}

func (r *Products) Delete(id uint) error {
	_, err := r.conn.Exec(
		context.Background(),
		`DELETE FROM products WHERE id=$1`,
		id,
	)

	return err
}
