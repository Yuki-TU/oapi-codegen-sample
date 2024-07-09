package handler

import (
	"context"

	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/usecase"
)

type GetAccountService interface {
	GetAccount(ctx context.Context) (usecase.GetAccountResponse, error)
}

type UpdateAccountUsecase interface {
	UpdateAccount(ctx context.Context, input gen.PutAccountJSONRequestBody) (usecase.UpdateAccountResponse, error)
}
