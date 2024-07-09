package controller

import (
	"github.com/Yuki-TU/oapi-codegen-sample/gen"
	"github.com/gin-gonic/gin"
)

type Controllers struct{}

func NewController() *Controllers {
	return &Controllers{}
}

func (c *Controllers) GetAccount(ctx *gin.Context, params gen.GetAccountRequestObject) (gen.GetAccountResponseObject, error) {
	return gen.GetAccount200JSONResponse{}, nil
}

func (c *Controllers) PutAccount(ctx *gin.Context, params gen.PutAccountRequestObject) (gen.PutAccountResponseObject, error) {
	return gen.PutAccount201JSONResponse{}, nil
}
