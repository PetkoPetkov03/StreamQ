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

-- name: FetchUserForSession :one
SELECT u.id, u.email, u.hash, u.username, u.profileid, p.userRole 
FROM USERS u
INNER JOIN PROFILES p ON p.id = u.profileid 
WHERE u.email = ? 
LIMIT 1;
