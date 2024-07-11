package usecase

import (
	"context"
	"database/sql"

	"github.com/Yuki-TU/oapi-codegen-sample/domain/model"
	"github.com/Yuki-TU/oapi-codegen-sample/myerrors"
	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	"github.com/cockroachdb/errors"
)

type GetAccount struct {
	Repo repository.Querier
	DB   repository.DBTX
}

func NewGetAccount(repo repository.Querier, db *sql.DB) *GetAccount {
	return &GetAccount{
		Repo: repo,
		DB:   db,
	}
}

type GetAccountResponse struct {
	UserID           model.UserID
	Email            string
	FamilyName       string
	FamilyNameKana   string
	FirstName        string
	FirstNameKana    string
	AcquisitionPoint int
	SendablePoint    int
}

func (ga *GetAccount) GetAccount(ctx context.Context) (GetAccountResponse, error) {
	userID := 1

	// Emailよりユーザ情報を取得する
	user, err := ga.Repo.GetByUserID(ctx, ga.DB, int64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return GetAccountResponse{}, errors.Join(err, myerrors.ErrNotUser)
		}
	}

	// 取得ポイントを取得する
	point, err := ga.Repo.GetPointByUserID(ctx, ga.DB, uint64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return GetAccountResponse{
				UserID:           model.UserID(user.ID),
				Email:            user.Email,
				AcquisitionPoint: 0,
			}, nil
		}
		return GetAccountResponse{}, errors.Wrap(err, "failed to get acquisition point")
	}

	return GetAccountResponse{
		UserID:           model.UserID(user.ID),
		Email:            user.Email,
		AcquisitionPoint: int(point.TransactionPoint),
	}, nil
}
