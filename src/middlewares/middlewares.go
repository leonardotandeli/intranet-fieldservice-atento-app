package middlewares

import (
	"app/src/cookies"
	"log"
	"net/http"
)

//Escreve informacoes no Terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

//Autentica e verifica a existencia de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := cookies.Ler(r); erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		// cookies.Ler
		proximaFuncao(w, r)
	}
}
