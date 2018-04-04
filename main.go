package main

import (
	"github.com/fatrbaby/marmot/crawler"
	"fmt"
)

func main() {
	body := crawler.FetchContentFrom("http://www.zhenai.com/zhenghun", true)

	fmt.Printf("%s", body)
}
