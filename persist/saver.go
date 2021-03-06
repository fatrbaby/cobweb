package persist

import (
	"context"
	"errors"
	"github.com/fatrbaby/cobweb/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		n := 0

		for {
			item := <-out
			log.Printf("Got item #%d: %v\n", n, item)
			if err := Save(client, index, item); err != nil {
				log.Printf("save item %v error: %v", item, err)
			}
			n++
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexer := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexer.Id(item.Id)
	}

	_, err := indexer.Do(context.Background())

	return err
}
