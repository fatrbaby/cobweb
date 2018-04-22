package command

import (
	"github.com/fatrbaby/cobweb/distributed/persist"
	"github.com/fatrbaby/cobweb/distributed/worker"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/scheduler"
	"github.com/urfave/cli"
	"strings"
)

func StartCrawl() cli.Command {
	command := cli.Command{
		Name: "crawl",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "index",
				Value: "dating_profile",
				Usage: "the index of storage",
			},
			cli.StringFlag{
				Name: "worker-ports",
				Value: ":7900,:7901",
				Usage: "the ports of worker rpc sever listening (comma separated)",
			},
		},
		Action: func(context *cli.Context) {
			//saver, err := persist.ItemSaver(context.String("index"))
			saver, err := persist.ItemSaver(":8700")

			if err != nil {
				panic(err)
			}

			pool := worker.CreateClientPool(strings.Split(context.String("worker-ports"), ","))
			processor := worker.CreateProcessor(pool)

			spider := engine.Spider{
				Url:    "http://www.zhenai.com/zhenghun",
				Parser: &parser.CityListParser{},
			}

			runner := engine.ConcurrentEngine{
				Scheduler:   &scheduler.QueuedScheduler{},
				WorkerCount: 50,
				ItemChannel: saver,
				Processor: processor,
			}

			runner.Run(spider)
		},
	}

	return command
}
