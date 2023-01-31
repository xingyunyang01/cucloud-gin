package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xingyunyang01/cucloud-gin/cgin"
	"github.com/xingyunyang01/cucloud-gin/cgin/task"
	"github.com/xingyunyang01/cucloud-gin/test/src/models"
	"github.com/xingyunyang01/cucloud-gin/test/src/service"
	"gorm.io/gorm"
)

//controllers层负责控制定义路由以及与service层进行整合

type ArticleClass struct {
	Db             *gorm.DB                `inject:"-"`
	ArticleService *service.ArticleService `inject:"-"`
}

func NewArticleClass() *ArticleClass {
	return &ArticleClass{}
}

func (this *ArticleClass) ArticleDetail(ctx *gin.Context) cgin.Model {
	news := models.NewArticleModel()
	cgin.Error(ctx.ShouldBindUri(news))
	cgin.Error(this.ArticleService.GetArticleDetail(news, this.Db))
	task.Task(this.UpdateViews, func() {
		this.UpdateViewsDone(news.NewsId)
	}, news.NewsId) //执行一个协程任务

	return news
}

func (this *ArticleClass) UpdateViews(params ...interface{}) {
	this.ArticleService.UpdateViews(params[0], this.Db)
}

func (this *ArticleClass) UpdateViewsDone(id int) {
	log.Println("点击量增加结束")
}

// 构造路由方法
func (this *ArticleClass) Build(cgin *cgin.Cgin) {
	cgin.Handle("GET", "/article/:id", this.ArticleDetail)
}

func (this *ArticleClass) Name() string {
	return "article"
}
