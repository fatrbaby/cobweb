package controller

import (
	"context"
	"github.com/fatrbaby/cobweb/engine"
	"github.com/fatrbaby/cobweb/web/model"
	"github.com/fatrbaby/cobweb/web/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
)

type SearchedResultHandler struct {
	view   view.SearchedResultView
	client *elastic.Client
}

func NewSearchedResultHandler(template string) SearchedResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchedResultHandler{client: client, view: view.CreateSearchedResultView(template)}
}

func (that SearchedResultHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	q := request.FormValue("q")

	from, err := strconv.Atoi(request.FormValue("from"))

	if err != nil {
		from = 0
	}

	var page model.SearchedResult
	page, err = that.searchFromElastic(q, from)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	err = that.view.Render(response, page)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}
}

func (that SearchedResultHandler) searchFromElastic(q string, from int) (model.SearchedResult, error) {
	var result model.SearchedResult
	result.Query = q
	q = rewriteQueryString(q)

	response, err := that.client.
		Search().
		Index("dating_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = response.TotalHits()
	result.Start = from
	result.Items = response.Each(reflect.TypeOf(engine.Item{}))

	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([\w]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
