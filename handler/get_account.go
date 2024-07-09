package handler

import (
	"fmt"

	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/myerrors"
	"github.com/cockroachdb/errors"
	"github.com/gin-gonic/gin"
)

type GetAccount struct {
	Service GetAccountService
}

func NewGetAccount(s GetAccountService) *GetAccount {
	return &GetAccount{Service: s}
}

func (gu *GetAccount) ServeHTTP(ctx *gin.Context) gen.GetAccountResponseObject {
	user, err := gu.Service.GetAccount(ctx)

	if err != nil {
		ctx.Error(err)
		if errors.Is(err, myerrors.ErrNotUser) {
			return gen.GetAccount401JSONResponse{
				N401UnauthorizedErrorJSONResponse: gen.N401UnauthorizedErrorJSONResponse{
					Message: "ユーザが存在しません。",
				},
			}
		}
		return gen.GetAccount500JSONResponse{
			N500ErrorJSONResponse: gen.N500ErrorJSONResponse{
				Message: "予期せぬエラーが発生しました。",
			},
		}
	}

	fmt.Println(user)
	return gen.GetAccount500JSONResponse{
		N500ErrorJSONResponse: gen.N500ErrorJSONResponse{
			Message: "予期せぬエラーが発生しました。",
		},
	}
	// return gen.GetAccount200JSONResponse{
	// 	UserId:           float32(user.UserID),
	// 	Email:            user.Email,
	// 	FamilyName:       user.FamilyName,
	// 	AcquisitionPoint: user.AcquisitionPoint,
	// 	FirstName:        user.FirstName,
	// 	FamilyNameKana:   user.FamilyNameKana,
	// 	FirstNameKana:    user.FirstNameKana,
	// 	SendablePoint:    user.SendablePoint,
	// }
}
