package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/xingyunyang01/cucloud-gin/cgin"
)

type TokenCheck struct {
}

func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

func (this *TokenCheck) OnRequest(ctx *gin.Context) error {
	if ctx.Query("token") == "" {
		cgin.Throw("token requered", 503, ctx)
	}
	return nil
}

func (this *TokenCheck) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
