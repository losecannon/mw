package render

import (
	"html/template"
	"net/http"
)

const (
	ext string = ".html"
	dir string = "templates/"
)

var (
	templates *template.Template
)

func HTML(rw http.ResponseWriter, name string, layout string, data interface{}) {
	templates = template.Must(template.ParseFiles(dir+name+ext, dir+layout+ext))
	err := templates.ExecuteTemplate(rw, "base", data)
	if err != nil {
		http.Error(rw, err.Error(), 500)
	}
}
