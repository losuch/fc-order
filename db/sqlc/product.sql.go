// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product (name, description, images_url, price, active, type_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id, name, description, images_url, price, active, type_id, created_at
`

type CreateProductParams struct {
	Name        string
	Description pgtype.Text
	ImagesUrl   pgtype.Text
	Price       pgtype.Numeric
	Active      pgtype.Numeric
	TypeID      int64
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Name,
		arg.Description,
		arg.ImagesUrl,
		arg.Price,
		arg.Active,
		arg.TypeID,
	)
	var i Product
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Description,
		&i.ImagesUrl,
		&i.Price,
		&i.Active,
		&i.TypeID,
		&i.CreatedAt,
	)
	return i, err
}

const getActiveProducts = `-- name: GetActiveProducts :many
SELECT product_id, name, description, images_url, price, active, type_id, created_at FROM product WHERE active = 1
`

func (q *Queries) GetActiveProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.Query(ctx, getActiveProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Description,
			&i.ImagesUrl,
			&i.Price,
			&i.Active,
			&i.TypeID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}