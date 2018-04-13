package main

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/scheduler"
)

func main() {
	/*
	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.CityListParser,
	}
	*/

	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun/chengdu",
		Parser: parser.CityParser,
	}

	runner := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	runner.Run(spider)
}
