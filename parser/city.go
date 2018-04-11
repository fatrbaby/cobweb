package parser

import (
	"github.com/fatrbaby/cobweb/engine"
	"regexp"
)

type City struct {
	Name []byte
	Link []byte
}

const (
	CityPattern     = `<a href="(http://album.zhenai.com/u/[0-9]+"[^>]*)>([^<]+)</a>`
	CityListPattern = `<a href="(http://www.zhenai.com/zhenghun/\w+)"[^>]*>([^<]+)</a>`
)

func CityParser(content []byte) engine.ParsedResult {
	re := regexp.MustCompile(CityPattern)

	matches := re.FindAllSubmatch(content, -1)
	var results = engine.ParsedResult{}

	for _, match := range matches {
		results.Items = append(results.Items, string(match[2]))
		results.Spiders = append(results.Spiders, engine.Spider{Url: string(match[1]), Parser: engine.NilParser})
	}

	return results
}

func CityListParser(content []byte) engine.ParsedResult {
	re := regexp.MustCompile(CityListPattern)

	matches := re.FindAllSubmatch(content, -1)
	var results = engine.ParsedResult{}

	for _, match := range matches {
		results.Items = append(results.Items, string(match[2]))
		results.Spiders = append(results.Spiders, engine.Spider{Url: string(match[1]), Parser: CityParser})
	}

	return results
}
