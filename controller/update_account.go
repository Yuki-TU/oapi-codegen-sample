package controller

import (
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/handler"
	"github.com/Yuki-TU/oapi-codegen-sample/repository"
	"github.com/Yuki-TU/oapi-codegen-sample/usecase"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) PutAccount(ctx *gin.Context, params gen.PutAccountRequestObject) (gen.PutAccountResponseObject, error) {
	// DIを行う
	repository := repository.New()
	uc := usecase.NewUpdateAcccount(repository, c.db)
	handler := handler.NewUpdatetAccount(uc)
	return handler.ServeHTTP(ctx, *params.Body), nil
}
