package engine

import (
	"github.com/fatrbaby/cobweb/crawler"
	"log"
)

func Worker(spider Spider) (ParsedResult, error) {
	body, err := crawler.Fetch(spider.Url, true)

	if err != nil {
		log.Printf("Fetch error on fetching: %s, %v\n", spider.Url, err)
		return ParsedResult{}, err
	}

	return spider.Parser.Parse(body, spider.Url), nil
}
