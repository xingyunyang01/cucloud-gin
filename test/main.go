package main

import (
	"log"

	"github.com/xingyunyang01/cucloud-gin/cgin"
	"github.com/xingyunyang01/cucloud-gin/test/src/configuration"
	"github.com/xingyunyang01/cucloud-gin/test/src/controllers"
	"github.com/xingyunyang01/cucloud-gin/test/src/middlewares"
)

func main() {
	cgin.Init().
		Config(configuration.NewDBConfig(), configuration.NewServiceConfig()).
		Attach(middlewares.NewTokenCheck()).
		Mount("v1", controllers.NewArticleClass()).
		Task("0/3 * * * * *", func() { //添加定时任务
			log.Println("开始定时任务")
		}).Launch()
}
