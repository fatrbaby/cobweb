package main

import (
	"github.com/fatrbaby/cobweb/command"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Cobweb - a website crawler"
	app.Usage = "cobweb [command] args..."
	app.Version = "0.5.0"

	app.Commands = []cli.Command{
		command.StartCrawl(),
		command.ServeWeb(),
		command.ServeRpcWorker(),
		command.ServeRpcSaver(),
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
