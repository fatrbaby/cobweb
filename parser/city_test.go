package parser

import (
	"fmt"
	"github.com/fatrbaby/marmot/crawler"
	"testing"
)

func TestCityParser(t *testing.T) {
	contents, err := crawler.Fetch("http://www.zhenai.com/zhenghun", true)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", contents)
}
