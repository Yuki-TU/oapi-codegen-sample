package usecase

import (
	"testing"

	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	mock_repository "github.com/Yuki-TU/oapi-codegen-sample/repository/_mock"
	"github.com/Yuki-TU/oapi-codegen-sample/repository/userrepo"
	mock_userrepo "github.com/Yuki-TU/oapi-codegen-sample/repository/userrepo/_mock"
	"github.com/Yuki-TU/oapi-codegen-sample/testutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccount(t *testing.T) {
	t.Parallel()
	type want struct {
		res UpdateAccountResponse
		err error
	}
	type getByUserID struct {
		userID int64
		res    userrepo.GetByUserIDRow
		err    error
		times  int
	}
	type updateAccount struct {
		params userrepo.UpdateUserParams
		err    error
		times  int
	}
	tests := map[string]struct {
		input gen.PutAccountJSONRequestBody
		getByUserID
		updateAccount
		want want
	}{
		"メールアドレスを更新する": {
			input: gen.PutAccountJSONRequestBody{
				FirstName:      "test",
				FamilyName:     "test",
				Email:          "test",
				FirstNameKana:  "test",
				FamilyNameKana: "test",
			},
			getByUserID: getByUserID{
				userID: 2,
				res: userrepo.GetByUserIDRow{
					ID:    2,
					Email: "test",
				},
				err:   nil,
				times: 1,
			},
			updateAccount: updateAccount{
				params: userrepo.UpdateUserParams{
					ID:             2,
					FirstName:      "test",
					FamilyName:     "test",
					Email:          "test",
					FirstNameKana:  "test",
					FamilyNameKana: "test",
				},
				err:   nil,
				times: 1,
			},
			want: want{
				res: UpdateAccountResponse{
					Status: "ok",
				},
				err: nil,
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			ctx := testutils.CreateContext(t)
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// モックの定義
			mockDB := mock_repository.NewMockDBer(ctrl)

			mockUserRepo := mock_userrepo.NewMockUserRepoer(ctrl)
			mockUserRepo.
				EXPECT().GetByUserID(ctx, mockDB, tt.getByUserID.userID).Return(tt.getByUserID.res, tt.getByUserID.err).Times(tt.getByUserID.times)

			// 実行
			ua := &GetAccount{
				UserRepo: mockUserRepo,
				DB:       mockDB,
			}
			got, err := ua.UpdateAccount(ctx, tt.input)

			// アサーション
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.res, got)
		})
	}
}
