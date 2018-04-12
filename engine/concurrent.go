package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Spider)
	SetMasterWorkerChannel(chan Spider)
}

func (ce *ConcurrentEngine) Run(spiders ...Spider) {
	for _, spider := range spiders {
		ce.Scheduler.Submit(spider)
	}

	in := make(chan Spider)
	out := make(chan ParsedResult)
	ce.Scheduler.SetMasterWorkerChannel(in)

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(in, out)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			fmt.Printf("Got item:%v\n", item)
		}

		for _, spider := range result.Spiders {
			ce.Scheduler.Submit(spider)
		}
	}
}

func createWorker(in chan Spider, out chan ParsedResult) {
	go func() {
		for {
			spider := <-in
			result, err := worker(spider)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
