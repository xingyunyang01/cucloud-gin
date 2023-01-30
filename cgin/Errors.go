package cgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_STATUS = "CGIN_STATUS"
)

// 捕获panic用的中间件
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
			}
		}()
		ctx.Next()
	}
}

// 抛出自定义的http error code
func Throw(err string, code int, context *gin.Context) {
	context.Set(HTTP_STATUS, code)
	panic(err)
}

// 对错误处理进行封装，可以抛出自定义错误信息的错误
func Error(err error, msg ...string) {
	if err == nil {
		return
	} else {
		errMsg := err.Error() //返回错误的string格式的内容
		if len(msg) > 0 {     //如果用户自定义了错误信息，则使用用户自定义的
			errMsg = msg[0]
		}
		panic(errMsg)
	}
}
