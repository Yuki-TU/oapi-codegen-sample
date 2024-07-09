package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Yuki-TU/oapi-codegen-sample/config"
	"github.com/Yuki-TU/oapi-codegen-sample/controller"
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error(fmt.Sprintf("failed to terminated server: %s", err.Error()))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	if err := config.Load(); err != nil {
		return err
	}

	if config.Get().Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if config.Get().Env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	// r := gin.New()
	// ミドルウェアの設定
	// r.Use(middleware.Recovery)
	// r.Use(middleware.Logger())
	// r.Use(middleware.CORSMiddleware())

	// DB関係初期化
	// db, cleanup, err := repository.NewDB(ctx, config.Get())
	// if err != nil {
	// 	return err
	// }
	// defer cleanup()

	db := repository.NewDB()
	defer db.Close()

	// ルーティング初期化
	impl := controller.NewController(db)
	strictServer := gen.NewStrictHandler(impl, nil)
	gen.RegisterHandlersWithOptions(
		r,
		strictServer,
		gen.GinServerOptions{BaseURL: "/api/"},
	)

	// サーバー起動
	log.Printf("Listening and serving HTTP on :%v", config.Get().Port)
	server := NewServer(r, fmt.Sprintf(":%d", config.Get().Port))
	return server.Run(ctx)
}
