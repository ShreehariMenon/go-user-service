-- name: CreateUser :execresult
INSERT INTO users (name, dob) VALUES (?, ?);

-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users SET name = ?, dob = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;