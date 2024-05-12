-- name: GetFilmList :many
SELECT * FROM film;

-- name: GetFilm :one
SELECT * FROM film WHERE film_id = $1;

-- name: CreateFilm :one
INSERT INTO film (name, yt_link, active, type) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateFilm :one
UPDATE film SET 
name = $1, 
yt_link = $2, 
active = $3, 
type = $4
WHERE film_id = $5 RETURNING *;

-- name: DeleteFilm :exec
DELETE FROM film WHERE film_id = $1;
