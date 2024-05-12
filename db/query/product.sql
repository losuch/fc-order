-- name: GetProductList :many
SELECT * FROM product;

-- name: GetProduct :one
SELECT * FROM product WHERE product_id = $1;

-- name: CreateProduct :one
INSERT INTO product (name, description, images_url, price, active, type_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateProduct :one
UPDATE product SET 
name = $1, 
description = $2,
images_url = $3, 
price = $4,
active = $5, 
type_id = $6
WHERE product_id = $7 RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE product_id = $1;
