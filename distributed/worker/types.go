package worker

import (
	"errors"
	"fmt"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/parser"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Spider struct {
	Url    string
	Parser SerializedParser
}

func SerializeSpider(spider engine.Spider) Spider {
	name, args := spider.Parser.Serialize()

	return Spider{
		Url: spider.Url,
		Parser: SerializedParser{
			name,
			args,
		},
	}
}

type ParsedResult struct {
	Items   []engine.Item
	Spiders []Spider
}

func SerializeReult(r engine.ParsedResult) ParsedResult {
	result := ParsedResult{Items: r.Items}

	for _, spider := range r.Spiders {
		result.Spiders = append(result.Spiders, SerializeSpider(spider))
	}

	return result
}

func DeserializeSpider(spider Spider) (engine.Spider, error) {
	p, err := DeserializeParser(spider.Parser)

	if err != nil {
		return engine.Spider{}, err
	}

	return engine.Spider{Url: spider.Url, Parser: p}, nil
}

func DeserializeResult(r ParsedResult) engine.ParsedResult  {
	result := engine.ParsedResult{Items: r.Items}

	for _, spider := range r.Spiders {
		s, err := DeserializeSpider(spider)

		if err != nil {
			log.Printf("unserializing spider error: %v\n", err)
			continue
		}

		result.Spiders = append(result.Spiders, s)
	}

	return result
}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case "ProfileParser":
		if username, ok := p.Args.(string); ok {
			return &parser.ProfileParser{Username: username}, nil
		} else {
			return nil, fmt.Errorf("invalid arg: %v", p.Args)
		}
	case "CityParser":
		return &parser.CityParser{}, nil
	case "CityListParser":
		return &parser.CityListParser{}, nil
	case "NilParser":
		return &engine.NilParser{}, nil
	default:
		return nil, errors.New("error parser")
	}
}
