package cgin

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type MiddlewareHandler struct {
	middlewares []Middleware
}

var middlewareHandler *MiddlewareHandler
var mid_once sync.Once

func getMiddlewareHandler() *MiddlewareHandler {
	mid_once.Do(func() {
		middlewareHandler = &MiddlewareHandler{}
	})
	return middlewareHandler
}

// 将用户添加的所有中间件函数都暂存到切片中
func (this *MiddlewareHandler) AddMiddleware(f ...Middleware) {
	this.middlewares = append(this.middlewares, f...)
}

// 执行中间件函数OnRequest
func (this *MiddlewareHandler) doRequest(ctx *gin.Context) {
	for _, f := range this.middlewares {
		if err := f.OnRequest(ctx); err != nil {
			Throw(err.Error(), 400, ctx)
		}
	}
}

// 执行中间件函数OnResponse
func (this *MiddlewareHandler) doResponse(ctx *gin.Context, ret interface{}) interface{} {
	var result = ret
	for _, f := range this.middlewares {
		r, err := f.OnResponse(ret)
		if err != nil {
			Throw(err.Error(), 400, ctx)
		}
		result = r
	}
	return result
}

func (this *MiddlewareHandler) RunMiddlerWare(responder Responder, ctx *gin.Context) interface{} {
	getMiddlewareHandler().doRequest(ctx)

	var ret interface{}
	if strfunc, ok := responder.(StringResponder); ok {
		ret = strfunc(ctx)
	} else if modelfunc, ok := responder.(ModelResponder); ok {
		ret = modelfunc(ctx)
	} else if modelsfunc, ok := responder.(ModelsResponder); ok {
		ret = modelsfunc(ctx)
	}

	return getMiddlewareHandler().doResponse(ctx, ret)
}
