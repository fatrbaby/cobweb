package view

import "path"

var viewPath string

func SetViewPath(p string)  {
	viewPath = p
}

func Load(view string) string {
	return path.Join(viewPath, view)
}