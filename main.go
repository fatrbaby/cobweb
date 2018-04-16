package main

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"github.com/fatrbaby/cobweb/persist"
	"github.com/fatrbaby/cobweb/scheduler"
	"github.com/fatrbaby/cobweb/web/controller"
	"github.com/urfave/cli"
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
			Name: "web",
			Usage: "run web server",
			Action: func(context *cli.Context) {
				web()
			},
		},
		{
			Name: "crawling",
			Usage: "crawling data from target",
			Action: func(context *cli.Context) {
				crawling()
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func crawling()  {
	saver, err := persist.ItemSaver("dating_profile")

	if err != nil {
		panic(err)
	}

	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.CityListParser,
	}

	runner := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 50,
		ItemChannel: saver,
	}

	runner.Run(spider)
}

func web() {
	http.Handle("/", http.FileServer(http.Dir("web/resources/assets")))
	http.Handle("/search", controller.NewSearchedResultHandler("web/resources/list.html"))
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
