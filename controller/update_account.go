package controller

import (
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/Yuki-TU/oapi-codegen-sample/handler"
	"github.com/Yuki-TU/oapi-codegen-sample/repository/userrepo"
	"github.com/Yuki-TU/oapi-codegen-sample/usecase"
	"github.com/gin-gonic/gin"
)

func (c *Controllers) PutAccount(ctx *gin.Context, params gen.PutAccountRequestObject) (gen.PutAccountResponseObject, error) {
	// DIを行う
	userrepo := userrepo.NewUserRepo()
	uc := usecase.NewUpdateAcccount(userrepo, c.db)
	handler := handler.NewUpdatetAccount(uc)
	return handler.ServeHTTP(ctx, *params.Body), nil
}
