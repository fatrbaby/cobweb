package command

import (
	"fmt"
	"github.com/fatrbaby/cobweb/distributed"
	"github.com/fatrbaby/cobweb/distributed/persist"
	"github.com/fatrbaby/cobweb/distributed/worker"
	"github.com/urfave/cli"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ServeRpcSaver() cli.Command {
	command := cli.Command{
		Name: "saver",
		Usage: "Start saver rpc server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 8700,
				Usage: "the port saver rpc server to listen",
			},
		},
		Action: func(context *cli.Context) {
			client, err := elastic.NewClient(elastic.SetSniff(false))

			if err != nil {
				panic(err)
			}

			port := context.Int("port")

			err = distributed.ServeRpc(fmt.Sprintf(":%d", port), &persist.ItemSaverService{
				Client: client,
				Index:  "dating_profile",
			})

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return command
}

func ServeRpcWorker() cli.Command {
	command := cli.Command{
		Name: "worker",
		Usage: "Start worker rpc server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 7900,
				Usage: "the port work rpc server to listen",
			},
		},
		Action: func(context *cli.Context) {
			port := context.Int("port")

			err := distributed.ServeRpc(fmt.Sprintf(":%d", port), &worker.CrawlService{})

			if err != nil {
				log.Fatal(err)
			}
		},
	}

	return command
}
