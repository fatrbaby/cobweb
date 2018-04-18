package engine

type Parser interface {
	Parse(contents []byte, url string) ParsedResult
	Serialize() (name string, args interface{})
}

type NilParser struct{}

func (np *NilParser) Parse(contents []byte, url string) ParsedResult {
	return ParsedResult{}
}

func (np *NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}
