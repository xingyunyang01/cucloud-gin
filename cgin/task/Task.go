package task

import "sync"

type TaskFunc func(params ...interface{}) //定于任务函数模板

var taskList chan *TaskExecutor //任务列表

type TaskExecutor struct {
	task     TaskFunc      //任务函数
	param    []interface{} //任务函数的参数
	callback func()        //用于结束任务函数的回调函数
}

// 构造函数
func NewTaskExecutor(task TaskFunc, param []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{task: task, param: param, callback: callback}
}

func (this *TaskExecutor) Exec() { //执行任务
	this.task(this.param...)
}

func init() {
	chlist := getTaskList() //初始化任务列表
	go func() {
		for t := range chlist { //只要通道里塞入任务函数，就会执行
			doTask(t)
		}
	}()
}

var once sync.Once //go语言提供的单例模式的库

func getTaskList() chan *TaskExecutor {
	once.Do(func() { //只执行一次
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

func doTask(t *TaskExecutor) {
	go func() {
		defer func() { //最后执行结束任务的回调函数
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec() //先执行任务函数
	}()
}

// 构建任务结构
func Task(task TaskFunc, cb func(), params ...interface{}) {
	if task == nil {
		return
	}
	go func() { //向队列里塞任务结构体
		getTaskList() <- NewTaskExecutor(task, params, cb)
	}()
}
