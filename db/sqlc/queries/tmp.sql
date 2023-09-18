-- name: GetAllTMP :many

SELECT id, title, created_at, is_deleted FROM tmp;

-- name: GetTMP :one

SELECT id, title, created_at, is_deleted FROM tmp WHERE id = $1;

-- name: GetTMPCount :one

SELECT COUNT(*) FROM tmp;