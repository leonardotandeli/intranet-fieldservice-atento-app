package rotas

import (
	"app/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Rota define a estrutura de uma rota.
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

//Configurar inicializa as rotas com o middleware.
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotaHome...)
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotasBaseDeConhecimento...)
	rotas = append(rotas, rotasCategoriasBaseDeConhecimento...)
	rotas = append(rotas, rotasClientesBaseDeConhecimento...)
	rotas = append(rotas, rotasMapaOperacional...)
	rotas = append(rotas, rotasSalaDeCrise...)
	rotas = append(rotas, rotasConsultas...)
	rotas = append(rotas, rotasContatos...)

	for _, rota := range rotas {

		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)

		} else {
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}

	}
	//fileServer define o caminho que ira servir os assets.
	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router

}
