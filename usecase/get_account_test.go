package usecase

import (
	"testing"

	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	mock_repository "github.com/Yuki-TU/oapi-codegen-sample/repository/_mock"
	"github.com/Yuki-TU/oapi-codegen-sample/testutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAccount(t *testing.T) {
	t.Parallel()
	type getByUserID struct {
		userID int64
		res    repository.GetByUserIDRow
		err    error
		times  int
	}
	type getPointByUserID struct {
		userID uint64
		res    repository.Transactions
		err    error
		times  int
	}
	type want struct {
		res GetAccountResponse
		err error
	}

	tests := map[string]struct {
		getByUserID
		getPointByUserID
		want want
	}{
		"メールアドレスを更新する": {
			getByUserID: getByUserID{
				userID: 1,
				res: repository.GetByUserIDRow{
					ID:    1,
					Email: "test",
				},
				err:   nil,
				times: 1,
			},
			getPointByUserID: getPointByUserID{
				userID: 1,
				res: repository.Transactions{
					ID:               1,
					TransactionPoint: 100,
				},
				err:   nil,
				times: 1,
			},
			want: want{
				res: GetAccountResponse{
					UserID:           1,
					Email:            "test",
					SendablePoint:    0,
					AcquisitionPoint: 100,
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
			mockDB := mock_repository.NewMockDBTX(ctrl)

			mockRepo := mock_repository.NewMockQuerier(ctrl)
			mockRepo.
				EXPECT().GetByUserID(ctx, mockDB, tt.getByUserID.userID).Return(tt.getByUserID.res, tt.getByUserID.err).Times(tt.getByUserID.times)

			mockRepo.
				EXPECT().GetPointByUserID(ctx, mockDB, tt.getPointByUserID.userID).Return(tt.getPointByUserID.res, tt.getPointByUserID.err).Times(tt.getPointByUserID.times)

			// 実行
			ua := &GetAccount{
				Repo: mockRepo,
				DB:   mockDB,
			}

			got, err := ua.GetAccount(ctx)

			// アサーション
			assert.ErrorIs(t, err, tt.want.err)
			assert.Equal(t, tt.want.res, got)
		})
	}
}
