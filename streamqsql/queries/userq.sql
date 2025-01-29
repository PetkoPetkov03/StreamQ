-- name: GetUserById :one
SELECT * FROM USERS
WHERE id = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO USERS (
  email, username, hash, profileid
) VALUES (
?, ?, ?, ?
)
RETURNING *;

-- name: CreateProfile :one
INSERT INTO PROFILES (role) VALUES(?) RETURNING *;
