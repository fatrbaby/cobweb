package parser

import (
	"github.com/fatrbaby/marmot/engine"
	"regexp"
)

type City struct {
	Name []byte
	Link []byte
}

const (
	CityListPattern = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`
)

func CityParser(content []byte) engine.ParsedResult {
	re := regexp.MustCompile(CityListPattern)

	matches := re.FindAllSubmatch(content, -1)
	var results = engine.ParsedResult{}

	for _, match := range matches {
		results.Items = append(results.Items, string(match[2]))
		results.Spiders = append(results.Spiders, engine.Spider{Url:string(match[1]), ParserFunc: engine.NilParser})
	}

	return results
}
