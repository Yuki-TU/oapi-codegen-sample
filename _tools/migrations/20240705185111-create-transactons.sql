
-- +migrate Up
CREATE TABLE `transactions` (
    `id`                 BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '取引の識別子',
    `sending_user_id`    BIGINT UNSIGNED NOT NULL COMMENT '送信ユーザのID',
    `receiving_user_id`  BIGINT UNSIGNED NOT NULL COMMENT '受信ユーザのID',
    `transaction_point`  INT NOT NULL COMMENT '取引ポイント',
    `transaction_at`     DATETIME(6) NOT NULL COMMENT '取引日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='取引';

-- +migrate Down
DROP TABLE `transactions`;
