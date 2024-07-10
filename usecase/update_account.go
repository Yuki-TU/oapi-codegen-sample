package usecase

import (
	"context"
	"database/sql"

	"github.com/Yuki-TU/oapi-codegen-sample/domain/model"
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/myerrors"
	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	"github.com/Yuki-TU/oapi-codegen-sample/repository/userrepo"
	"github.com/cockroachdb/errors"
)

type UpdateAccount struct {
	// UserRepo repository.UserRepoer
	UserRepo userrepo.Querier
	Tx       repository.Beginner
}

func NewUpdateAcccount(userrepo userrepo.Querier, tx repository.Beginner) *UpdateAccount {
	return &UpdateAccount{
		UserRepo: userrepo,
		Tx:       tx,
	}
}

type UpdateAccountResponse struct {
	Status string
}

func (ua *UpdateAccount) UpdateAccount(ctx context.Context, input gen.PutAccountJSONRequestBody) (UpdateAccountResponse, error) {
	tx, err := ua.Tx.BeginTx(ctx, nil)
	if err != nil {
		return UpdateAccountResponse{}, errors.Wrap(err, "failed to begin transaction")
	}
	defer func() { _ = tx.Rollback() }()

	// 一旦固定
	userID := model.UserID(2)

	// ユーザー存在確認
	user, err := ua.UserRepo.GetByUserID(ctx, tx, int64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return UpdateAccountResponse{}, errors.Join(err, myerrors.ErrNotUser)
		}
		return UpdateAccountResponse{}, errors.Wrap(err, "failed to get user")
	}

	// 更新
	err = ua.UserRepo.UpdateUser(ctx, tx, userrepo.UpdateUserParams{
		ID:             user.ID,
		FirstName:      input.FirstName,
		FamilyName:     input.FamilyName,
		Email:          input.Email,
		FirstNameKana:  input.FirstNameKana,
		FamilyNameKana: input.FamilyNameKana,
	})
	if err != nil {
		return UpdateAccountResponse{}, errors.Wrap(err, "failed to update user")
	}
	if err := tx.Commit(); err != nil {
		return UpdateAccountResponse{}, errors.Wrap(err, "failed to commit transaction")
	}

	return UpdateAccountResponse{Status: "ok"}, nil
}
