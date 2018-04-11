package main

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
)

func main() {
	spider := engine.Spider{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.CityListParser,
	}

	engine.Run(spider)
}
