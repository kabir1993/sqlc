
-- name: ListAll :many
SELECT users.username , users.pass, authors.name from users join authors on authors.id = users.userid;

-- name: GetAuthor :many
SELECT * FROM authors;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = $1;

-- name: GetUser :many
SELECT * FROM users;

-- name: UpdateAuthor :one
UPDATE authors SET bio = $2
WHERE id = $1
RETURNING *;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  username, pass
) VALUES (
  $1, $2
)RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE userid = $1;