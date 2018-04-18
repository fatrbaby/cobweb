package persist

import (
	"github.com/fatrbaby/cobweb/distributed"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (is *ItemSaverService)Save(item engine.Item, result *string) error {
	err := persist.Save(is.Client, is.Index, item)

	if err == nil {
		*result = "success"
	}

	return err
}

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := distributed.NewClient(host)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		n := 0

		for {
			item := <-out
			log.Printf("Got item #%d: %v\n", n, item)
			var result string
			if err := client.Call("ItemSaverService.Save", item, &result); err != nil {
				log.Printf("save item %v error: %v", item, err)
			}
			n++
		}
	}()

	return out, nil
}