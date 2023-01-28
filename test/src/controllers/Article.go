package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xingyunyang01/cucloud-gin/cgin"
	"github.com/xingyunyang01/cucloud-gin/cgin/orm"
	"github.com/xingyunyang01/cucloud-gin/cgin/task"
	"github.com/xingyunyang01/cucloud-gin/test/src/models"
	"gorm.io/gorm"
)

type ArticleClass struct {
	*orm.GormAdapter
}

func NewArticleClass() *ArticleClass {
	return &ArticleClass{}
}

func (this *ArticleClass) ArticleDetail(ctx *gin.Context) cgin.Model {
	news := models.NewArticleModel()
	cgin.Error(ctx.ShouldBindUri(news))
	cgin.Error(this.Table("mynews").Where("id=?", news.NewsId).Find(news).Error)

	task.Task(this.UpdateViews, func() {
		this.UpdateViewsDone(news.NewsId)
	}, news.NewsId) //执行一个协程任务

	return news
}

func (this *ArticleClass) UpdateViews(params ...interface{}) {
	this.Table("mynews").Where("id=?", params[0]).Update("views", gorm.Expr("views+1"))
}

func (this *ArticleClass) UpdateViewsDone(id int) {
	log.Println("点击量增加结束")
}

// 构造路由方法
func (this *ArticleClass) Build(cgin *cgin.Cgin) {
	cgin.Handle("GET", "/article/:id", this.ArticleDetail)
}
