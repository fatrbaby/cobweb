package command

import (
	"github.com/fatrbaby/cobweb/distributed/persist"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/scheduler"
	"github.com/urfave/cli"
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
		},
		Action: func(context *cli.Context) {
			//saver, err := persist.ItemSaver(context.String("index"))
			saver, err := persist.ItemSaver(":8700")

			if err != nil {
				panic(err)
			}

			spider := engine.Spider{
				Url:    "http://www.zhenai.com/zhenghun",
				Parser: &parser.CityListParser{},
			}

			runner := engine.ConcurrentEngine{
				Scheduler:   &scheduler.QueuedScheduler{},
				WorkerCount: 50,
				ItemChannel: saver,
			}

			runner.Run(spider)
		},
	}

	return command
}
