package engine

type Spider struct {
	Url    string
	Parser Parser
}

type Item struct {
	Id      string
	Url     string
	Type    string
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

type ParsedResult struct {
	Spiders []Spider
	Items   []Item
}
