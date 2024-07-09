package userrepo

import (
	"net/http/httptest"
	"testing"

	"github.com/Yuki-TU/oapi-codegen-sample/repository/userrepo/testdata"
	"github.com/Yuki-TU/oapi-codegen-sample/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRepository_users_UpdateEmail(t *testing.T) {
	type want struct {
		email string
		err   error
	}
	type input struct {
		email string
		id    int64
	}

	tests := map[string]struct {
		input input
		want  want
	}{
		"ユーザーID_1のメールを更新する": {
			input: input{
				email: "before@sample.com",
				id:    1,
			},
			want: want{
				email: "after@sample.com",
				err:   nil,
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run(n, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			tx, err := testutils.OpenDBForTest(t).BeginTxx(ctx, nil)
			assert.NoError(t, err)
			t.Cleanup(func() {
				_ = tx.Rollback()
			})

			testdata.Users(t, ctx, tx, func(users []*testdata.User) {

			})

			// 実行
			r := &UserRepo{}
			err = r.UpdateUserMail(ctx, tx,
				UpdateUserMailParams{
					ID:    tt.input.id,
					Email: tt.input.email,
				},
			)
			assert.Nil(t, err)

			// 確認
			// err = r.UpdateUserMail(ctx, tx, UpdateUserMailParams{
			// 	Email: tt.want.email,
			// 	ID:    tt.input.id,
			// })
			assert.Nil(t, err)
			got, err := r.GetByUserID(ctx, tx, tt.input.id)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.email, got.Email)
		})
	}
}
