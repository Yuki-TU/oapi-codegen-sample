package transactionrepo

import (
	"context"
)

// UserRepoer は、ユーザ情報に関するリポジトリのインターフェース
type TransactionRepoer interface {
	GetByUserID(ctx context.Context, db DBTX, receivingUserID uint64) (Transactions, error)
}

type TransactionsRepo struct{}

func NewUserRepo() *TransactionsRepo {
	return &TransactionsRepo{}
}

func (ur *TransactionsRepo) GetByUserID(ctx context.Context, db DBTX, receivingUserID uint64) (Transactions, error) {
	return New(db).GetByUserID(ctx, receivingUserID)
}

// UserRepoer が UserRepo を実装していることを確認
var _ TransactionRepoer = (*TransactionsRepo)(nil)
