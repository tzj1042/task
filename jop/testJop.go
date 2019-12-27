package jop

import "fmt"

type TestJob struct {
}

func (this TestJob) Run() {
	fmt.Println("任务一逻辑...")
}
