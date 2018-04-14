package engine

type Spider struct {
	Url    string
	Parser Parser
}

type ParsedResult struct {
	Spiders []Spider
	Items   []Item
}

type Parser func(contents []byte, url string) ParsedResult

type Item struct {
	Id string
	Url string
	Type string
	Payload interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Spider)
	WorkerChannel() chan Spider
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Spider)
}

func NilParser([]byte) ParsedResult {
	return ParsedResult{}
}
