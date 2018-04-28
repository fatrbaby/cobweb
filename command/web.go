package command

import (
	"fmt"
	"github.com/fatrbaby/cobweb/web/controller"
	"github.com/urfave/cli"
	"log"
	"net/http"
)

func ServeWeb() cli.Command {
	command := cli.Command{
		Name: "web",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 8090,
				Usage: "the port of web server listening",
			},
		},
		Action: func(context *cli.Context) {
			port := context.Int("port")

			staticFilesHandler := http.FileServer(http.Dir("web/resources/assets"))
			http.Handle("/assets/", http.StripPrefix("/assets/", staticFilesHandler))
			http.Handle("/", controller.NewHomeHandler("web/resources/home.html"))
			http.Handle("/search", controller.NewSearchedResultHandler("web/resources/list.html"))

			log.Printf("Serve on %d", port)

			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
		},
	}

	return command
}
