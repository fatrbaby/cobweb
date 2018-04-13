package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (ce *ConcurrentEngine) Run(spiders ...Spider) {
	ce.Scheduler.Run()
	for _, spider := range spiders {
		ce.Scheduler.Submit(spider)
	}

	out := make(chan ParsedResult)

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChannel(), out, ce.Scheduler)
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

func createWorker(in chan Spider, out chan ParsedResult,  notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			spider := <-in
			result, err := worker(spider)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
