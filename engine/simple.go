package engine

import (
	"github.com/fatrbaby/cobweb/crawler"
	"log"
)

type SimpleEngine struct {
}

func (se *SimpleEngine) Run(spiders ...Spider) {
	var seeds []Spider

	for _, spider := range spiders {
		seeds = append(seeds, spider)
	}

	for len(seeds) > 0 {
		seed := seeds[0]
		seeds = seeds[1:]

		result, err := worker(seed)

		if err != nil {
			continue
		}

		seeds = append(seeds, result.Spiders...)

		for _, item := range result.Items {
			log.Printf("Got url: %v", item)
		}
	}
}

func worker(spider Spider) (ParsedResult, error) {
	log.Printf("Fetching: %s\n", spider.Url)
	body, err := crawler.Fetch(spider.Url, true)

	if err != nil {
		log.Printf("Fetch error on fetching: %s, %v\n", spider.Url, err)
		return ParsedResult{}, err
	}

	return spider.Parser(body), nil
}
