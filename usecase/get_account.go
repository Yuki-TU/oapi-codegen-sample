package usecase

import (
	"context"
	"database/sql"

	"github.com/Yuki-TU/oapi-codegen-sample/domain/model"
	"github.com/Yuki-TU/oapi-codegen-sample/myerrors"
	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	"github.com/Yuki-TU/oapi-codegen-sample/repository/transactionrepo"
	"github.com/Yuki-TU/oapi-codegen-sample/repository/userrepo"
	"github.com/cockroachdb/errors"
)

// type GetAccount struct {
// 	// UserRepo userrepo.Querier
// 	UserRepo        userrepo.DBer
// 	TransactionRepo transactionrepo.Querier
// 	Tx              repository.Beginner
// }

// func NewGetAccount(urepo userrepo.DBer, trepo transactionrepo.Querier, tx *sql.DB) *GetAccount {
// 	return &GetAccount{
// 		UserRepo:        urepo,
// 		TransactionRepo: trepo,
// 		Tx:              tx,
// 	}
// }

// type GetAccountResponse struct {
// 	UserID           model.UserID
// 	Email            string
// 	FamilyName       string
// 	FamilyNameKana   string
// 	FirstName        string
// 	FirstNameKana    string
// 	AcquisitionPoint int
// 	SendablePoint    int
// }

// func (ga *GetAccount) GetAccount(ctx context.Context) (GetAccountResponse, error) {
// 	tx, err := ga.Tx.BeginTx(ctx, nil)
// 	if err != nil {
// 		return GetAccountResponse{}, errors.Wrap(err, "failed to begin transaction")
// 	}
// 	defer func() { _ = tx.Rollback() }()
// 	qtx := ga.UserRepo.WithTx(tx)

// 	userID := 2
// 	// Emailよりユーザ情報を取得する
// 	// user, err := ga.UserRepo.GetByUserID(ctx, int64(userID))
// 	user, err := qtx.GetByUserID(ctx, int64(userID))
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return GetAccountResponse{}, errors.Join(err, myerrors.ErrNotUser)
// 		}
// 	}
// 	err = qtx.UpdateUserMail(ctx, userrepo.UpdateUserMailParams{
// 		// err = ga.UserRepo.UpdateUserMail(ctx, userrepo.UpdateUserMailParams{
// 		Email: "hoge@tin.com",
// 		ID:    int64(userID),
// 	})
// 	if err == nil {
// 		return GetAccountResponse{}, errors.New("failed to update user email")
// 		// return GetAccountResponse{}, errors.Wrap(err, "failed to update user email")
// 	}

// 	// 取得ポイントを取得する
// 	// userID := utils.GetUserID(ctx)
// 	point, err := ga.TransactionRepo.GetByUserID(ctx, uint64(userID))
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			if err := tx.Commit(); err != nil {
// 				return GetAccountResponse{}, errors.Wrap(err, "failed to commit transaction")
// 			}
// 			return GetAccountResponse{
// 				UserID: model.UserID(user.ID),
// 				Email:  user.Email,
// 			}, nil

// 		}
// 		return GetAccountResponse{}, errors.Wrap(err, "failed to get acquisition point")
// 	}

// 	if err := tx.Commit(); err != nil {
// 		return GetAccountResponse{}, errors.Wrap(err, "failed to commit transaction")
// 	}

// 	return GetAccountResponse{
// 		UserID:           model.UserID(user.ID),
// 		Email:            user.Email,
// 		AcquisitionPoint: int(point.TransactionPoint),
// 	}, nil
// }
// *******************

type GetAccount struct {
	// UserRepo userrepo.Querier
	UserRepo        userrepo.UserRepoer
	TransactionRepo transactionrepo.Querier
	Tx              repository.Beginner
}

func NewGetAccount(urepo userrepo.UserRepoer, trepo transactionrepo.Querier, tx *sql.DB) *GetAccount {
	return &GetAccount{
		UserRepo:        urepo,
		TransactionRepo: trepo,
		Tx:              tx,
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
	tx, err := ga.Tx.BeginTx(ctx, nil)
	if err != nil {
		return GetAccountResponse{}, errors.Wrap(err, "failed to begin transaction")
	}
	defer func() { _ = tx.Rollback() }()

	userID := 1
	// Emailよりユーザ情報を取得する
	// user, err := ga.UserRepo.GetByUserID(ctx, int64(userID))
	user, err := ga.UserRepo.GetByUserID(ctx, tx, int64(userID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return GetAccountResponse{}, errors.Join(err, myerrors.ErrNotUser)
		}
	}
	err = ga.UserRepo.UpdateUserMail(ctx, tx, userrepo.UpdateUserMailParams{
		// err = ga.UserRepo.UpdateUserMail(ctx, userrepo.UpdateUserMailParams{
		Email: "hoge@tin.comhoge",
		ID:    int64(userID),
	})
	if err != nil {
		return GetAccountResponse{}, errors.New("failed to update user email")
		// return GetAccountResponse{}, errors.Wrap(err, "failed to update user email")
	}

	// 取得ポイントを取得する
	// userID := utils.GetUserID(ctx)
	// point, err := ga.TransactionRepo.GetByUserID(ctx, uint64(userID))
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		if err := tx.Commit(); err != nil {
	// 			return GetAccountResponse{}, errors.Wrap(err, "failed to commit transaction")
	// 		}
	// 		return GetAccountResponse{
	// 			UserID: model.UserID(user.ID),
	// 			Email:  user.Email,
	// 		}, nil

	// 	}
	// 	return GetAccountResponse{}, errors.Wrap(err, "failed to get acquisition point")
	// }

	if err := tx.Commit(); err != nil {
		return GetAccountResponse{}, errors.Wrap(err, "failed to commit transaction")
	}

	return GetAccountResponse{
		UserID: model.UserID(user.ID),
		Email:  user.Email,
		// AcquisitionPoint: int(point.TransactionPoint),
	}, nil
}
