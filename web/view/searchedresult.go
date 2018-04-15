package view

import (
	"github.com/fatrbaby/cobweb/web/model"
	"html/template"
	"io"
)

type SearchedResultView struct {
	template *template.Template
}

func CreateSearchedResultView(name string) SearchedResultView {
	return SearchedResultView{
		template: template.Must(template.ParseFiles(name)),
	}
}

func (that SearchedResultView) Render(writer io.Writer, data model.SearchedResult) error {
	return that.template.Execute(writer, data)
}
