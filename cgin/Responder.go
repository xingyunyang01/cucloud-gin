package cgin

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 定义一个Responder的切片
var ResponderList []Responder

// 初始化切片，塞入StringResponder(new之后是指针类型)
func init() {
	ResponderList = []Responder{new(StringResponder),
		new(JsonResponder),
		new(VoidResponder),
	}
}

// 这一个类就是为了凑不同类型的HandlerFunc的，因为在goft.Handle函数里面的this.g.Handle中，需要这个参数。
type Responder interface {
	RespondTo() gin.HandlerFunc
}

// 该函数用来判断handler的类型是否和ResponderList切片中的某一种类型一样，如果一样，就通过反射执行RespondTo。
func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, responder := range ResponderList { //遍历ResponderList
		r_ref := reflect.ValueOf(responder).Elem()    //由于ResponderList里放的指针，因此需要使用Elem取出值。
		if h_ref.Type().ConvertibleTo(r_ref.Type()) { //判断h_ref是否能转换成r_ref的类型
			//如果可以，将handler的value赋给r_ref，此步骤相当于var i int = 10
			//效果是StringResponder = UserList
			r_ref.Set(h_ref)
			//这一步是先r_ref.Interface().(Responder)取出原始数据的值，之后调用RespondTo方法
			return r_ref.Interface().(Responder).RespondTo()
		}
	}

	return nil
}

// 返回值为string类型的HandlerFunc
type StringResponder func(ctx *gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, getMiddlewareHandler().RunMiddlerWare(this, ctx).(string))
	}
}

// 返回值为Json类型的HandlerFunc
type Json interface{}
type JsonResponder func(ctx *gin.Context) Json

func (this JsonResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, getMiddlewareHandler().RunMiddlerWare(this, ctx))
	}
}

// 返回值为void类型的HandlerFunc
type Void struct{}
type VoidResponder func(ctx *gin.Context) Void

func (this VoidResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		getMiddlewareHandler().RunMiddlerWare(this, ctx)
	}
}
