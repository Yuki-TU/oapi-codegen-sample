package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/Yuki-TU/oapi-codegen-sample/config"
	"github.com/go-sql-driver/mysql"
)

var (
	instance *sql.DB
	once     sync.Once
)

// DBインスタンスを取得する
// シングルトンパターンで行うためのやり方
func NewDB() *sql.DB {
	var err error
	once.Do(func() {
		instance, err = connect()
		if err != nil {
			panic(err)
		}
	})
	return instance
}

// DB接続
func connect() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 config.Get().DBUser,
		Passwd:               config.Get().DBPassword,
		Addr:                 fmt.Sprintf("%s:%d", config.Get().DBHost, config.Get().DBPort),
		DBName:               config.Get().DBName,
		ParseTime:            true,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(1 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

// トランザクションを開始するためのインターフェース
type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// Beginnerインターフェースがsql.DBのメソッドを定義しているかのチェック
var _ Beginner = (*sql.DB)(nil)
