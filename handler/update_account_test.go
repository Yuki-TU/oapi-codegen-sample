package handler

import (
	"testing"

	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	mock_handler "github.com/Yuki-TU/oapi-codegen-sample/handler/_mock"
	"github.com/Yuki-TU/oapi-codegen-sample/myerrors"
	"github.com/Yuki-TU/oapi-codegen-sample/testutils"
	"github.com/Yuki-TU/oapi-codegen-sample/usecase"
	"github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServeHTTP(t *testing.T) {
	t.Parallel()
	type want struct {
		res gen.PutAccountResponseObject
	}
	type updateAccountUsecase struct {
		input gen.PutAccountJSONRequestBody
		res   usecase.UpdateAccountResponse
		err   error
		times int
	}

	tests := map[string]struct {
		input gen.PutAccountJSONRequestBody
		updateAccountUsecase
		want want
	}{
		"正常系-ユーザー情報を元に更新し、okメッセージと共にステータス200を返す": {
			input: gen.PutAccountJSONRequestBody{
				FirstName:      "test",
				FamilyName:     "test",
				Email:          "test@sample.com",
				FirstNameKana:  "test",
				FamilyNameKana: "test",
			},
			updateAccountUsecase: updateAccountUsecase{
				input: gen.PutAccountJSONRequestBody{
					FirstName:      "test",
					FamilyName:     "test",
					Email:          "test@sample.com",
					FirstNameKana:  "test",
					FamilyNameKana: "test",
				},
				res: usecase.UpdateAccountResponse{
					Status: "ok",
				},
				err:   nil,
				times: 1,
			},
			want: want{
				res: gen.PutAccount201JSONResponse{
					Status: "ok",
				},
			},
		},
		"メールアドレスが正しくない場合は、400エラーを返す": {
			input: gen.PutAccountJSONRequestBody{
				FirstName:      "test",
				FamilyName:     "test",
				Email:          "test",
				FirstNameKana:  "test",
				FamilyNameKana: "test",
			},
			updateAccountUsecase: updateAccountUsecase{
				input: gen.PutAccountJSONRequestBody{},
				res:   usecase.UpdateAccountResponse{},
				err:   nil,
				times: 0,
			},
			want: want{
				res: gen.PutAccount400JSONResponse{
					N400ErrorJSONResponse: gen.N400ErrorJSONResponse{
						Message: "email: メールアドレスの形式が正しくありません.",
					},
				},
			},
		},
		"名前が51の場合は、400エラーを返す": {
			input: gen.PutAccountJSONRequestBody{
				FirstName:      "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
				FamilyName:     "test",
				Email:          "test@sample.com",
				FirstNameKana:  "test",
				FamilyNameKana: "test",
			},
			updateAccountUsecase: updateAccountUsecase{
				input: gen.PutAccountJSONRequestBody{},
				res:   usecase.UpdateAccountResponse{},
				err:   nil,
				times: 0,
			},
			want: want{
				res: gen.PutAccount400JSONResponse{
					N400ErrorJSONResponse: gen.N400ErrorJSONResponse{
						Message: "firstName: 名前は50文字以内で入力してください.",
					},
				},
			},
		},
		"ユーザーが存在しない場合は、404エラーを返す": {
			input: gen.PutAccountJSONRequestBody{
				FirstName:      "test",
				FamilyName:     "test",
				Email:          "test@sample.com",
				FirstNameKana:  "test",
				FamilyNameKana: "test",
			},
			updateAccountUsecase: updateAccountUsecase{
				input: gen.PutAccountJSONRequestBody{
					FirstName:      "test",
					FamilyName:     "test",
					Email:          "test@sample.com",
					FirstNameKana:  "test",
					FamilyNameKana: "test",
				},
				res:   usecase.UpdateAccountResponse{},
				err:   errors.Wrap(myerrors.ErrNotUser, "failed to get user"),
				times: 1,
			},
			want: want{
				res: gen.PutAccount404JSONResponse{
					N404ErrorJSONResponse: gen.N404ErrorJSONResponse{
						Message: myerrors.ClientErrNoUser,
					},
				},
			},
		},
		"予期せぬエラーは、500エラーを返す": {
			input: gen.PutAccountJSONRequestBody{
				FirstName:      "test",
				FamilyName:     "test",
				Email:          "test@sample.com",
				FirstNameKana:  "test",
				FamilyNameKana: "test",
			},
			updateAccountUsecase: updateAccountUsecase{
				input: gen.PutAccountJSONRequestBody{
					FirstName:      "test",
					FamilyName:     "test",
					Email:          "test@sample.com",
					FirstNameKana:  "test",
					FamilyNameKana: "test",
				},
				res:   usecase.UpdateAccountResponse{},
				err:   myerrors.ErrUnexpected,
				times: 1,
			},
			want: want{
				res: gen.PutAccount500JSONResponse{
					N500ErrorJSONResponse: gen.N500ErrorJSONResponse{
						Message: myerrors.ClientErrInternalServer,
					},
				},
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ctx := testutils.CreateGinContext(t)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// モックの定義
			mockUpdateAccountUsecase := mock_handler.NewMockUpdateAccountUsecase(ctrl)
			mockUpdateAccountUsecase.EXPECT().UpdateAccount(ctx, tt.input).Return(tt.res, tt.err).Times(tt.times)

			// 実行
			got := NewUpdatetAccount(mockUpdateAccountUsecase).ServeHTTP(ctx, tt.input)

			// アサーション
			assert.Equal(t, tt.want.res, got)
		})
	}
}
