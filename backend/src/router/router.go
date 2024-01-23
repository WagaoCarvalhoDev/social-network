package router

import "github.com/gorilla/mux"

// Retorna um router com as rotas configuradas
func NewAPIRouter() *mux.Router {
	return mux.NewRouter()
}
