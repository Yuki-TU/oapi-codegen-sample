-- name: GetByUserID :one
SELECT
  `id`,
  `email`
FROM
  `users`
WHERE `id` = ?;
