package engine

import (
	"github.com/fatrbaby/cobweb/entity"
)

var visited = make(map[string]bool)

type Processor func(Spider) (ParsedResult, error)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChannel chan Item
	Processor   Processor
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
		ce.CreateWorker(ce.Scheduler.WorkerChannel(), out, ce.Scheduler)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			if _, ok := item.Payload.(entity.Profile); ok {
				go func() { ce.ItemChannel <- item }()
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

func (ce ConcurrentEngine)CreateWorker(in chan Spider, out chan ParsedResult, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			spider := <-in
			result, err := ce.Processor(spider)

			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
