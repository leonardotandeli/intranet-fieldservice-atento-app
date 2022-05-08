package router

import (
	"app/src/controllers"
	"app/src/router/rotas"
	"net/http"

	"github.com/gorilla/mux"
)

//Retorna um router com as rotas configuradas.
// Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(controllers.Handler404)

	return rotas.Configurar(r)
}
