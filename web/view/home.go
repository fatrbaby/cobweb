package view

import (
	"html/template"
	"io"
)

type HomeView struct {
	template *template.Template
}

func (that HomeView) Render(writer io.Writer) error {
	return that.template.Execute(writer, nil)
}

func CreateHomeView(name string) HomeView {
	return HomeView{
		template: template.Must(template.ParseFiles(name)),
	}
}
