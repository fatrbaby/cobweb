package command

import (
	"fmt"
	"github.com/fatrbaby/cobweb/web/controller"
	"github.com/fatrbaby/cobweb/web/view"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"path/filepath"
)

func ServeWeb() cli.Command {
	command := cli.Command{
		Name: "web",
		Usage: "Start web server",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "port",
				Value: 8090,
				Usage: "the port of web server listening",
			},
		},
		Action: func(context *cli.Context) {
			port := context.Int("port")
			abs, err := filepath.Abs("./../web/resources/views")

			if err != nil {
				log.Fatal("can not ger abs path of views")
			}

			view.SetViewPath(abs)

			staticDir, _ := filepath.Abs("./../web/resources/assets")

			staticFilesHandler := http.FileServer(http.Dir(staticDir))

			http.Handle("/assets/", http.StripPrefix("/assets/", staticFilesHandler))
			http.Handle("/", controller.NewHomeHandler("home.html"))
			http.Handle("/search", controller.NewSearchedResultHandler("list.html"))

			log.Printf("Serve on %d", port)

			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
		},
	}

	return command
}
