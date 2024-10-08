// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateUsers(ctx context.Context, db DBTX, arg CreateUsersParams) (sql.Result, error)
	GetByUserID(ctx context.Context, db DBTX, id int64) (GetByUserIDRow, error)
	GetPointByUserID(ctx context.Context, db DBTX, receivingUserID uint64) (Transactions, error)
	UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) error
	UpdateUserMail(ctx context.Context, db DBTX, arg UpdateUserMailParams) error
}

var _ Querier = (*Queries)(nil)
