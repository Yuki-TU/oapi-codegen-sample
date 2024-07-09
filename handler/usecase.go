package handler

import (
	"github.com/gin-gonic/gin"
)

type GetAccountService interface {
	GetAccount(ctx *gin.Context) (service.GetAccountResponse, error)
}
