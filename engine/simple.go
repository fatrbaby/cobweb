package engine

import (
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

		result, err := Worker(seed)

		if err != nil {
			continue
		}

		seeds = append(seeds, result.Spiders...)

		for _, item := range result.Items {
			log.Printf("Got url: %v", item)
		}
	}
}
