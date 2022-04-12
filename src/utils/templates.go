package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

//Carrega os arquivos HTML na variavel templates
func CarregarTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

//Executar template, renderiza a página html
func ExecutarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, template, dados)

}
