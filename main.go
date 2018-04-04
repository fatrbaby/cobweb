package main

import (
	"fmt"
	"github.com/fatrbaby/marmot/crawler"
	"github.com/fatrbaby/marmot/parser"
)

func main() {
	body := crawler.FetchContentFrom("http://www.zhenai.com/zhenghun", true)
	cities := parser.FindCities(body)

	for _, city := range cities {
		fmt.Printf("%s: %s\n", city.Name, city.Link)
	}
}
