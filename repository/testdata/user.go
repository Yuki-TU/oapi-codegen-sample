package testdata

import (
	"context"
	"testing"
	"time"

	"github.com/Yuki-TU/oapi-codegen-sample/utils/clock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type User struct {
	ID             int64     `json:"id" db:"id"`                             // ユーザーの識別子
	FamilyName     string    `json:"family_name" db:"family_name"`           // 苗字
	FamilyNameKana string    `json:"family_name_kana" db:"family_name_kana"` // 苗字カナ
	FirstName      string    `json:"first_name" db:"first_name"`             // 名前
	FirstNameKana  string    `json:"first_name_kana" db:"first_name_kana"`   // 名前カナ
	Email          string    `json:"email" db:"email"`                       // メールアドレス
	Password       string    `json:"password" db:"password"`                 // パスワードハッシュ
	SendingPoint   int       `json:"sending_point" db:"sending_point"`       // 送信可能ポイント
	CreatedAt      time.Time `json:"created_at" db:"created_at"`             // レコード作成日時
	UpdateAt       time.Time `json:"update_at" db:"update_at"`               // レコード修正日時
}

// ユーザーテーブルにデータを挿入する
func Users(t *testing.T, ctx context.Context, tx *sqlx.Tx, setter func(users []*User)) []*User {
	t.Helper()
	// ユーザ一覧を削除する
	if _, err := tx.ExecContext(ctx, `DELETE FROM delete_users;`); err != nil {
		assert.Error(t, err)
	}
	if _, err := tx.ExecContext(ctx, `DELETE FROM users;`); err != nil {
		assert.Error(t, err)
	}

	c := clock.FixedClocker{
		IsAsia: true,
	}

	// ユーザーデータのデフォルト値
	users := []*User{
		{
			FirstName:      "太郎",
			FirstNameKana:  "タロウ",
			FamilyName:     "本田",
			FamilyNameKana: "ホンダ",
			Email:          "honda.taro@sample.com",
			Password:       "honda.pass",
			SendingPoint:   0,
			CreatedAt:      c.Now(),
			UpdateAt:       c.Now(),
		},
		{
			FirstName:      "葵",
			FirstNameKana:  "あおい",
			FamilyName:     "斉藤",
			FamilyNameKana: "さいとう",
			Email:          "saito.aoi@example.com",
			Password:       "aoi.pass",
			SendingPoint:   100,
			CreatedAt:      c.Now(),
			UpdateAt:       c.Now(),
		},
		{
			FirstName:      "拓也",
			FirstNameKana:  "たくや",
			FamilyName:     "木村",
			FamilyNameKana: "きむら",
			Email:          "kimura.takuya@example.com",
			Password:       "kimura.pass",
			SendingPoint:   800,
			CreatedAt:      c.Now(),
			UpdateAt:       c.Now(),
		},
	}

	// データの上書き
	setter(users)

	// データ道入力
	result, err := tx.NamedExecContext(ctx, `
		INSERT INTO users
			(first_name, first_name_kana, family_name, family_name_kana, email, password, sending_point, created_at, update_at)
		VALUES
			(:first_name, :first_name_kana, :family_name, :family_name_kana, :email, :password, :sending_point, :created_at, :update_at);`, users)
	assert.NoError(t, err)

	// 自動発行されたユーザーIDを取得
	id, err := result.LastInsertId()
	assert.NoError(t, err)

	users[0].ID = id
	users[1].ID = id + 1
	users[2].ID = id + 2

	return users
}
