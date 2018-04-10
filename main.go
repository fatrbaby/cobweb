package main

import (
	"github.com/fatrbaby/marmot/engine"
	"github.com/fatrbaby/marmot/parser"
)

func main() {
	spider := engine.Spider{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.CityParser,
	}

	engine.Run(spider)
}
