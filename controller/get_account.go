package controller

import (
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/handler"
	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	"github.com/Yuki-TU/oapi-codegen-sample/usecase"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) GetAccount(ctx *gin.Context, params gen.GetAccountRequestObject) (gen.GetAccountResponseObject, error) {
	// DIを行う
	repository := repository.New()
	uc := usecase.NewGetAccount(repository, c.db)
	handler := handler.NewGetAccount(uc)

	return handler.ServeHTTP(ctx), nil
}
