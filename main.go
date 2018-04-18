package main

import (
	"github.com/fatrbaby/cobweb/distributed"
	Persist "github.com/fatrbaby/cobweb/distributed/persist"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/persist"
	"github.com/fatrbaby/cobweb/scheduler"
	"github.com/fatrbaby/cobweb/web/controller"
	"github.com/urfave/cli"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"net/http"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Cobweb - a website crawler"
	app.Usage = "cobweb [command] args..."
	app.Version = "0.5.0"

	app.Commands = []cli.Command{
		{
			Name:   "web",
			Usage:  "run web server",
			Action: commandWeb,
		},
		{
			Name:   "crawl",
			Usage:  "crawling data from target",
			Action: commandCrawling,
		},
		{
			Name:   "rpc",
			Usage:  "start rpc server",
			Action: commandRpc,
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func commandCrawling(_ *cli.Context) {
	saver, err := persist.ItemSaver("dating_profile")
	// saver, err := Persist.ItemSaver(":8700")

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
}

func commandWeb(_ *cli.Context) {
	http.Handle("/", http.FileServer(http.Dir("web/resources/assets")))
	http.Handle("/search", controller.NewSearchedResultHandler("web/resources/list.html"))
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}

func commandRpc(_ *cli.Context) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	err = distributed.ServeRpc(":8700", Persist.ItemSaverService{
		Client: client,
		Index:  "dating_profile",
	})

	if err != nil {
		log.Fatal(err)
	}
}
