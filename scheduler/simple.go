package scheduler

import "github.com/fatrbaby/cobweb/engine"

type SimpleScheduler struct {
	workerChannel chan engine.Spider
}

func (se *SimpleScheduler) WorkerChannel() chan engine.Spider {
	return se.workerChannel
}

func (se *SimpleScheduler) WorkerReady(chan engine.Spider) {
}

func (se *SimpleScheduler) Run() {
	se.workerChannel = make(chan engine.Spider)
}

func (se *SimpleScheduler) Submit(spider engine.Spider) {
	go func() { se.workerChannel <- spider }()
}

