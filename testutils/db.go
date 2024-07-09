package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// // テストにおいてデータベース接続する
// // 実際にデーターベース接続する
// func OpenDBForTest(t *testing.T) *sql.DB {
// 	t.Helper()

// 	host := "db"
// 	if _, defined := os.LookupEnv("CI"); defined {
// 		host = "127.0.0.1"
// 	}

// 	db, err := sql.Open(
// 		"mysql",
// 		fmt.Sprintf("admin:password@tcp(%s:3306)/point_app_test?parseTime=true&loc=Asia%%2FTokyo", host),
// 	)
// 	assert.NoError(t, err, "cannot open db")

// 	t.Cleanup(
// 		func() { _ = db.Close() },
// 	)

// 	// AUTO_INCREMENTをリセットする
// 	_, err = db.Exec("ALTER TABLE `users` AUTO_INCREMENT = 1;")
// 	assert.NoError(t, err)

// 	return db
// }

// テストにおいてデータベース接続する
// 実際にデーターベース接続する
func OpenDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	host := "db"
	if _, defined := os.LookupEnv("CI"); defined {
		host = "127.0.0.1"
	}

	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("admin:password@tcp(%s:3306)/point_app_test?parseTime=true&loc=Asia%%2FTokyo", host),
	)
	assert.NoError(t, err, "cannot open db")

	t.Cleanup(
		func() { _ = db.Close() },
	)

	xdb := sqlx.NewDb(db, "mysql")

	// AUTO_INCREMENTをリセットする
	_, err = xdb.Exec("ALTER TABLE `users` AUTO_INCREMENT = 1;")
	assert.NoError(t, err)

	return xdb
}

func NewTxForMock(t *testing.T, ctx context.Context) *sql.Tx {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	t.Cleanup(
		func() { _ = db.Close() },
	)

	if err := db.Ping(); err != nil {
		assert.Error(t, err)
	}

	// トランザクションの開始、コミットをモックする
	// 正常系を返す
	mock.ExpectBegin()
	mock.ExpectCommit()

	mockTx, err := db.BeginTx(ctx, nil)
	if err != nil {
		assert.Error(t, err)
	}
	return mockTx
}
