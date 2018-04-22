package worker

import (
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/imooc-crawler/crawler_distributed/rpcsupport"
	"net/rpc"
)

func CreateProcessor(clients chan *rpc.Client) engine.Processor {
	return func(spider engine.Spider) (engine.ParsedResult, error) {
		s := SerializeSpider(spider)
		var r ParsedResult

		client := <-clients
		err := client.Call("CrawlService.Process", s, &r)

		if err != nil {
			return engine.ParsedResult{}, err
		}

		return DeserializeResult(r), nil
	}
}

func CreateClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, host := range hosts {
		client, err := rpcsupport.NewClient(host)

		if err != nil {
			continue
		}

		clients = append(clients, client)
	}

	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
