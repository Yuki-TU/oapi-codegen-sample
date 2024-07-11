-- name: GetByUserID :one
SELECT
  `id`,
  `email`
FROM
  `users`
WHERE `id` = ?;


-- name: UpdateUserMail :exec
UPDATE
  `users`
SET
  `email` = ?
WHERE `id` = ?;


-- name: UpdateUser :exec
UPDATE
  `users`
SET
  `first_name` = ?,
  `family_name` = ?,
  `first_name_kana` = ?,
  `family_name_kana` = ?,
  `email` = ?
WHERE `id` = ?;

-- name: CreateUsers :execresult
INSERT INTO `users` (
  `family_name`, `family_name_kana`, `first_name`, `first_name_kana`, `email`, `password`, `sending_point`, `created_at`, `update_at`
) 
VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?
);
