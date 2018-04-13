package scheduler

import (
	"github.com/fatrbaby/cobweb/engine"
)

type QueuedScheduler struct {
	spiderChannel chan engine.Spider
	workerChannel chan chan engine.Spider
}

func (qs *QueuedScheduler) Submit(spider engine.Spider) {
	qs.spiderChannel <- spider
}

func (qs *QueuedScheduler) SetMasterWorkerChannel(chan engine.Spider) {
	panic("implement me")
}

func (qs *QueuedScheduler) WorkerReady(spider chan engine.Spider) {
	qs.workerChannel <- spider
}

func (qs *QueuedScheduler) Run() {
	qs.spiderChannel = make(chan engine.Spider)
	qs.workerChannel = make(chan chan engine.Spider)

	go func() {
		var spiderQueue []engine.Spider
		var workerQueue []chan engine.Spider

		for {
			var activeSpider engine.Spider
			var activeWorker chan engine.Spider

			if len(spiderQueue) > 0 && len(workerQueue) > 0 {
				activeSpider = spiderQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
			case spider := <-qs.spiderChannel:
				spiderQueue = append(spiderQueue, spider)
			case worker := <-qs.workerChannel:
				workerQueue = append(workerQueue, worker)
			case activeWorker <- activeSpider:
				spiderQueue = spiderQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
