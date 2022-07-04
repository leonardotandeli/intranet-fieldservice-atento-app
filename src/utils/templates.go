package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//Carrega os arquivos HTML na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("tmpl/*.html"))
	templates = template.Must(templates.ParseGlob("tmpl/includes/*.html"))
}

//ExecutarTemplate renderiza a página html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)

}
