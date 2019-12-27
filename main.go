package main

import (
	"febsTask/jop"
	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

func init() {
	logs.NewLogger()
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["error", "warning", "info", "debug"]}`)
	logs.SetLogger(logs.AdapterConsole)
	logs.Debug("初始化log")
}

func main() {
	i := 0
	c := cron.New()

	//AddFunc
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		//logs.Info(fmt.Sprintf("cron running:%v",i))
		logs.Info("cron running：", i)
	})

	//AddJob方法
	c.AddJob(spec, jop.TestJob{})

	//启动计划任务
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}
