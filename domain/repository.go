package domain

import (
	"context"
)

// Userに対するインターフェース
type UserRepo interface {
	FindUserByEmail(ctx context.Context, db repository.Queryer, e string, columns ...string) (entities.User, error)
	GetUserByID(ctx context.Context, db repository.Queryer, ID model.UserID) (entities.User, error)
	DeleteUserByID(ctx context.Context, db repository.Execer, ID model.UserID) (int64, error)
	RegisterUser(ctx context.Context, db repository.Execer, u *entities.User) error
	UpdatePassword(ctx context.Context, db repository.Execer, email, pass *string) error
	UpdateEmail(ctx context.Context, db repository.Execer, userID model.UserID, newEmail string) error
	UpdateAccount(ctx context.Context, db repository.Execer, email, familyName, familyNameKana, firstName, firstNameKana *string) error
	GetAll(ctx context.Context, db repository.Queryer, columns ...string) ([]*entities.User, error)
}
