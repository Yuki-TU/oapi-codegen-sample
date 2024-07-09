package testutils

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func CreateContext(t *testing.T) context.Context {
	t.Helper()
	ctx := context.Background()
	return ctx
}

func CreateGinContext(t *testing.T) *gin.Context {
	t.Helper()
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	return ctx
}
