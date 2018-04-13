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
	Submit(Spider)
	SetMasterWorkerChannel(chan Spider)
	WorkerReady(chan Spider)
	Run()
}

func NilParser([]byte) ParsedResult {
	return ParsedResult{}
}
