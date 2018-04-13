package engine

import (
	"fmt"
	"github.com/fatrbaby/cobweb/entity"
)

var visited = make(map[string]bool)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (ce *ConcurrentEngine) Run(spiders ...Spider) {
	ce.Scheduler.Run()
	for _, spider := range spiders {
		if isDuplicate(spider.Url) {
			continue
		}

		ce.Scheduler.Submit(spider)
	}

	out := make(chan ParsedResult)

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.WorkerChannel(), out, ce.Scheduler)
	}

	personCount := 0

	for {
		result := <-out

		for _, item := range result.Items {
			if _, ok := item.(entity.Profile); ok {
				personCount++
				fmt.Printf("Got item [%d]:%v\n", personCount, item)
			}
		}

		for _, spider := range result.Spiders {
			if isDuplicate(spider.Url) {
				continue
			}

			ce.Scheduler.Submit(spider)
		}
	}
}

func isDuplicate(url string) bool {
	if visited[url] {
		return true
	}

	visited[url] = true

	return false
}

func createWorker(in chan Spider, out chan ParsedResult, notifier ReadyNotifier) {
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
