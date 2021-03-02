package commons

import (
	"html/template"
)

func GetTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.html"))
}
