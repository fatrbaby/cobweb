package command

import (
	"fmt"
	"github.com/fatrbaby/cobweb/distributed/persist"
	"github.com/fatrbaby/cobweb/distributed/worker"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/scheduler"
	"github.com/urfave/cli"
	"strings"
)

func Dispatch() cli.Command {
	command := cli.Command{
		Name: "dispatch",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name: "saver-port",
				Value: 8700,
				Usage: "the port of saver rpc server listening",
			},
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
			saver, err := persist.ItemSaver(fmt.Sprintf(":%d", context.Int("saver-port")))

			if err != nil {
				panic(err)
			}

			pool := worker.CreateClientPool(parseHosts(context.String("worker-ports")))
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

func parseHosts(hostsConfig string) []string  {
	hosts := strings.Split(hostsConfig, ",")

	for i, host := range hosts {
		if !strings.HasPrefix(host, ":") {
			hosts[i] = ":" + host
		}
	}

	return hosts
}
