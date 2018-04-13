package main

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/scheduler"
)

func main() {
	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.CityListParser,
	}

	runner := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	runner.Run(spider)
}
