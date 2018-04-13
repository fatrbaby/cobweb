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
	CityPattern     = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
	CityListPattern = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)"[^>]*>([^<]+)</a>`
)

var (
	CityMatcher = regexp.MustCompile(CityPattern)
	CityListMatcher = regexp.MustCompile(CityListPattern)
)

func CityParser(content []byte) engine.ParsedResult {
	matches := CityMatcher.FindAllSubmatch(content, -1)
	var results = engine.ParsedResult{}

	for _, match := range matches {
		name := string(match[2])
		results.Items = append(results.Items, name)
		parser := func(c []byte) engine.ParsedResult {
			return ProfileParser(c, name)
		}
		results.Spiders = append(results.Spiders, engine.Spider{Url: string(match[1]), Parser: parser})
	}

	return results
}

func CityListParser(content []byte) engine.ParsedResult {
	matches := CityListMatcher.FindAllSubmatch(content, -1)
	var results = engine.ParsedResult{}

	for _, match := range matches {
		results.Items = append(results.Items, string(match[2]))
		results.Spiders = append(results.Spiders, engine.Spider{Url: string(match[1]), Parser: CityParser})
	}

	return results
}
