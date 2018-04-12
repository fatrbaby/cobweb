package scheduler

import "github.com/fatrbaby/cobweb/engine"

type SimpleScheduler struct {
	WorkerChannel chan engine.Spider
}

func (se *SimpleScheduler) Submit(spider engine.Spider) {
	go func() { se.WorkerChannel <- spider }()
}

func (se *SimpleScheduler) SetMasterWorkerChannel(c chan engine.Spider) {
	se.WorkerChannel = c
}
