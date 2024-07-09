package handler

import (
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/myerrors"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type updateAccount struct {
	Usecase UpdateAccountUsecase
}

func NewUpdatetAccount(s UpdateAccountUsecase) *updateAccount {
	return &updateAccount{Usecase: s}
}

func (gu *updateAccount) ServeHTTP(ctx *gin.Context, input gen.PutAccountJSONRequestBody) gen.PutAccountResponseObject {
	// バリデーション
	err := validation.ValidateStruct(&input,
		validation.Field(
			&input.Email,
			is.Email.Error("メールアドレスの形式が正しくありません"),
			validation.Required.Error("メールアドレスは必須です"),
		),
		validation.Field(
			&input.FamilyName,
			validation.Length(1, 50).Error("名字は50文字以内で入力してください"),
			is.UTFLetter.Error("名字は全角文字で入力してください"),
		),
		validation.Field(
			&input.FirstName,
			validation.Length(1, 50).Error("名前は50文字以内で入力してください"),
			is.UTFLetter.Error("名前は全角文字で入力してください"),
		),
		validation.Field(
			&input.FamilyNameKana,
			validation.Length(1, 50).Error("名字は50文字以内で入力してください"),
			is.UTFLetter.Error("名字は全角文字で入力してください"),
		),
		validation.Field(
			&input.FirstNameKana,
			validation.Length(1, 50).Error("名前は50文字以内で入力してください"),
			is.UTFLetter.Error("名前は全角文字で入力してください"),
		),
	)
	if err != nil {
		ctx.Error(err)
		return gen.PutAccount400JSONResponse{
			N400ErrorJSONResponse: gen.N400ErrorJSONResponse{
				Message: err.Error(),
			},
		}
	}

	_, err = gu.Usecase.UpdateAccount(ctx, input)
	if err != nil {
		ctx.Error(err)
		if errors.Is(err, myerrors.ErrNotUser) {
			return gen.PutAccount404JSONResponse{
				N404ErrorJSONResponse: gen.N404ErrorJSONResponse{
					Message: myerrors.ClientErrNoUser,
				},
			}
		}
		return gen.PutAccount500JSONResponse{
			N500ErrorJSONResponse: gen.N500ErrorJSONResponse{
				Message: myerrors.ClientErrInternalServer,
			},
		}
	}
	return gen.PutAccount201JSONResponse{
		Status: "ok",
	}
}
