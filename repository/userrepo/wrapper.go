package userrepo

import (
	"context"
)

// UserRepoer は、ユーザ情報に関するリポジトリのインターフェース
type UserRepoer interface {
	GetByUserID(ctx context.Context, db DBTX, id int64) (GetByUserIDRow, error)
	UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) error
	UpdateUserMail(ctx context.Context, db DBTX, arg UpdateUserMailParams) error
}

type UserRepo struct{}

// NewUserRepo は、UserRepoのポインタを返す
func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

// GetByUserID は、ユーザ情報を取得する
func (ur *UserRepo) GetByUserID(ctx context.Context, db DBTX, id int64) (GetByUserIDRow, error) {
	return New(db).GetByUserID(ctx, id)
}

// UpdateUser は、ユーザ情報を更新する
func (ur *UserRepo) UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) error {
	return New(db).UpdateUser(ctx, arg)
}

// UpdateUserMail は、ユーザのメールアドレスを更新する
func (ur *UserRepo) UpdateUserMail(ctx context.Context, db DBTX, arg UpdateUserMailParams) error {
	return New(db).UpdateUserMail(ctx, arg)
}

// UserRepoer が UserRepo を実装していることを確認
var _ UserRepoer = (*UserRepo)(nil)
