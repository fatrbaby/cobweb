package main

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/persist"
	"github.com/fatrbaby/cobweb/scheduler"
)

func main() {
	saver, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}
	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.CityListParser,
	}
	/*
	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun/chengdu",
		Parser: parser.CityParser,
	}
	*/

	runner := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChannel: saver,
	}

	runner.Run(spider)
}
