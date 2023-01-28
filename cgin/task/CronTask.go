package task

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var onceCron sync.Once

var cronTask *cron.Cron //定时任务

// 初始化定时任务(单例模式)
func GetCronTask() *cron.Cron {
	onceCron.Do(func() {
		cronTask = cron.New(cron.WithSeconds())
	})

	return cronTask
}
