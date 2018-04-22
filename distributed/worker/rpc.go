package worker

import "github.com/fatrbaby/cobweb/engine"

type CrawlService struct{}

func (cs *CrawlService) Process(spider Spider, result *ParsedResult) error {
	engineSpider, err := DeserializeSpider(spider)

	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineSpider)

	if err != nil {
		return err
	}

	*result = SerializeReult(engineResult)

	return nil
}