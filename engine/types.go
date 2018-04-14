package engine

type Spider struct {
	Url    string
	Parser func([]byte) ParsedResult
}

type ParsedResult struct {
	Spiders []Spider
	Items   []Item
}

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
