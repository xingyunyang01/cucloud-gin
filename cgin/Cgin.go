package cgin

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xingyunyang01/cucloud-gin/cgin/fairing"
	"github.com/xingyunyang01/cucloud-gin/cgin/ioc"
	"github.com/xingyunyang01/cucloud-gin/cgin/task"
)

type Cgin struct {
	*gin.Engine                   //gin.New()的返回值
	g            *gin.RouterGroup //路由组
	currentGroup string           // temp-var for group string
	exprData     map[string]interface{}
}

func Init() *Cgin {
	g := &Cgin{Engine: gin.New(), exprData: make(map[string]interface{})}
	g.Use(ErrorHandler())             //强迫加载的异常处理中间件
	ioc.BeanFactory.Set(InitConfig()) //整个配置加载进bean中

	return g
}

func (this *Cgin) Launch() {
	var port int32 = 8080
	if config := ioc.BeanFactory.Get((*SysConfig)(nil)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	task.GetCronTask().Start()
	this.Run(fmt.Sprintf(":%d", port))
}

func (this *Cgin) Mount(group string, classes ...IClass) *Cgin {
	this.g = this.Group(group)
	for _, class := range classes {
		this.currentGroup = group
		class.Build(this)
		this.DBBeans(class) //初始化控制器实体类中的数据库连接对象句柄
	}
	return this
}

// 该方法现在的调用者是gin.Engine。
// 重载Handle方法，在该方法里调用this.g.Handle，由于this.g是gin.RouterGroup，且在Mount方法中已经获得了组名
// 因此重载该方法后，该方法的效果变成了调用组的Handle方法。这样index,user实体类的build去调用handle时，就可以实现组的效果。

// handle的改造：
// 1.将传入的handlers列表，改为只传一个空类型的handler
// 2.将handler断言成实体类的方法类型，如果断言成功，则可以将该函数塞到ctx.String里面了。
func (this *Cgin) Handle(httpMethod, relativePath string, handler interface{}) *Cgin {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}

	return this
}

type Bean interface {
	Name() string
}

// 设定数据库连接对象
func (this *Cgin) DBBeans(dbbeans ...Bean) *Cgin {
	for _, bean := range dbbeans {
		this.exprData[bean.Name()] = bean
		ioc.BeanFactory.Set(bean)
	}
	return this
}

// 中间件构造方法
// 将中间件实现方法封装到了用户中间件类，并通过接口的方式实现了统一
func (this *Cgin) Attach(f fairing.Fairing) *Cgin {
	this.Use(func(ctx *gin.Context) { //添加中间件函数
		err := f.OnRequest(ctx) //具体的中间件业务逻辑
		if err != nil {         //报错则返回错误
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else { //不报错则通过中间件，继续向下走
			ctx.Next()
		}
	})
	return this
}

func (this *Cgin) Task(expr string, f func()) *Cgin {
	_, err := task.GetCronTask().AddFunc(expr, f)
	if err != nil {
		log.Println(err)
	}
	return this
}
