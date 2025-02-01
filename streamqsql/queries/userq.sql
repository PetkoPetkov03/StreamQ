-- name: GetUserById :one
SELECT * FROM USERS
WHERE id = ? LIMIT 1;

-- name: CheckIfEmailExists :one
SELECT COUNT(*) FROM USERS WHERE email = ?;

-- name: CreateUser :one
INSERT INTO USERS (
  email, username, hash, profileid
) VALUES (
?, ?, ?, ?
)
RETURNING id;

-- name: CreateProfile :one
INSERT INTO PROFILES (userRole) VALUES(?) RETURNING id;