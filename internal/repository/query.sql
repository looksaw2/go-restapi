-- name: GetMovie :one
SELECT * FROM movies 
WHERE id = $1 LIMIT 1;  

-- name: GetListMovie :many
SELECT * FROM movies
ORDER BY id;


-- name: CreateMovie :one
INSERT INTO movies (
    name ,
    email 
)VALUES(
    $1,$2
)RETURNING *;


-- name: UpdateMovie :one
UPDATE movies
SET name = $2, email = $3
WHERE id = $1
RETURNING *;


-- name: DeleteMovie :exec
DELETE FROM movies
WHERE id = $1;
