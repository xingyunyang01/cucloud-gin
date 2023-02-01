package cgin

import "github.com/gin-gonic/gin"

//用来规范中间件代码和功能的接口
type Middleware interface {
	OnRequest(*gin.Context) error
	OnResponse(result interface{}) (interface{}, error)
}
