package engine

type Spider struct {
	Url    string
	Parser func([]byte) ParsedResult
}

type ParsedResult struct {
	Spiders []Spider
	Items   []interface{}
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
