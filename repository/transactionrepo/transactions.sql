-- name: GetByUserID :one
SELECT
  `id`,
  `sending_user_id`,
  `receiving_user_id`,
  `transaction_point`,
  `transaction_at`
FROM
  `transactions`
WHERE `receiving_user_id` = ?;
