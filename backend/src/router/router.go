package router

import (
	"backend/src/router/routers"

	"github.com/gorilla/mux"
)

// Retorna um router com as rotas configuradas
func NewAPIRouter() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigRouter(r)
}
