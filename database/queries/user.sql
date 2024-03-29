-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  photo,
  verified,
  password,
  role
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;




-- name: CreateOrUpdateUser :one
WITH updated_user AS (
  UPDATE users
  SET 
    name = COALESCE(sqlc.narg('name')::text, name),
    photo = COALESCE(sqlc.narg('photo')::text, photo),
    verified = COALESCE(sqlc.narg('verified')::bool, verified),
    password = COALESCE(sqlc.narg('password')::text, password),
    role = COALESCE(sqlc.narg('role')::text, role),
    updated_at = CURRENT_TIMESTAMP
  WHERE users.email = $2
  RETURNING id, name, email, photo, verified, password, role, created_at, updated_at
),
inserted_user AS (
  INSERT INTO users (name, email, photo, verified, password, role)
  SELECT $1, $2, $3, $4, $5, $6
  WHERE NOT EXISTS (SELECT 1 FROM updated_user)
  RETURNING id, name, email, photo, verified, password, role, created_at, updated_at
)
SELECT * FROM updated_user
UNION ALL
SELECT * FROM inserted_user;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: IsUserEmailExist :one
SELECT EXISTS ( SELECT 1 FROM users WHERE email = $1);

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
set name = COALESCE(NULLIF($2, ''), name),
email = COALESCE(NULLIF($3, ''), email),
password = COALESCE(NULLIF($4, ''), password),
photo = COALESCE(NULLIF($5, ''), photo),
updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
