package main

import (
	"github.com/fatrbaby/cobweb/web/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("resources/assets")))
	http.Handle("/search", controller.NewSearchedResultHandler("resources/list.html"))
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
