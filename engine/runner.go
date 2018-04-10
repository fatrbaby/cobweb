package engine

import (
	"github.com/fatrbaby/marmot/crawler"
	"log"
)

func Run(spiders ...Spider) {
	var seeds []Spider

	for _, spider := range spiders {
		seeds = append(seeds, spider)
	}

	for len(seeds) > 0 {
		seed := seeds[0]
		seeds = seeds[1:]
		body, err := crawler.Fetch(seed.Url, true)

		if err != nil {
			log.Printf("Fetch error on fetching: %s, %v", seed.Url, err)
			continue
		}

		result := seed.Parser(body)
		seeds = append(seeds, result.Spiders...)

		for _, item := range result.Items {
			log.Printf("Got url: %v", item)
		}
	}
}
