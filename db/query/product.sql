-- name: GetActiveProducts :many
SELECT * FROM product WHERE active = 1;

-- name: CreateProduct :one
INSERT INTO product (name, description, images_url, price, active, type_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateOrderStatus :one
UPDATE product SET active = $1 WHERE id = $2 RETURNING *;