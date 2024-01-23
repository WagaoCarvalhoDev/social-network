package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RouterApi struct {
	Uri                    string
	Method                 string
	HandlerFunction        func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func ConfigRouter(r *mux.Router) *mux.Router {
	// routers contém a definição das rotas da API.
	routers := routerUsers

	// Itera sobre as rotas definidas e adiciona cada uma ao roteador.
	for _, router := range routers {
		r.HandleFunc(router.Uri, router.HandlerFunction).Methods(router.Method)
	}

	// Retorna o roteador atualizado com as rotas da API.
	return r
}
