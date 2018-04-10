package engine

type Spider struct {
	Url string
	ParserFunc func([]byte) ParsedResult
}

type ParsedResult struct {
	Spiders []Spider
	Items []interface{}
}

func NilParser([]byte) ParsedResult  {
	return ParsedResult{}
}
