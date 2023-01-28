package main

import (
	"log"

	"github.com/xingyunyang01/cucloud-gin/cgin"
	"github.com/xingyunyang01/cucloud-gin/cgin/orm"
	"github.com/xingyunyang01/cucloud-gin/test/src/controllers"
)

func main() {
	cgin.Init().
		DBBeans(orm.NewGormAdapter()).
		Mount("v1", controllers.NewArticleClass()).
		Task("0/3 * * * * *", func() { //添加定时任务
			log.Println("开始定时任务")
		}).Launch()
}
